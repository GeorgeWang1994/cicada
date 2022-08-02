package rpc

import (
	"context"
	"errors"
	"fmt"
	"github.com/GeorgeWang1994/cicada/module/ev/cc"
	"github.com/GeorgeWang1994/cicada/module/ev/gg"
	"github.com/GeorgeWang1994/cicada/module/ev/sender/queue"
	"github.com/GeorgeWang1994/cicada/module/ev/store/db"
	"github.com/GeorgeWang1994/cicada/module/pkg/model"
	pb "github.com/GeorgeWang1994/cicada/module/proto/api/ev"
	"github.com/garyburd/redigo/redis"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Event struct {
	pb.UnimplementedEventServiceServer
}

func New() *Event {
	return &Event{}
}

func (t *Event) Ping(ctx context.Context, request *pb.Empty) (*pb.Response, error) {
	return &pb.Response{}, nil
}

// ReceiveEvent 接收来自探针的事件
func (t *Event) ReceiveEvent(ctx context.Context, request *pb.ReceiveEventRequest) (*pb.Response, error) {
	if len(request.Events) <= 0 {
		return &pb.Response{}, errors.New("empty events")
	}
	// 如果没有kafka，启动多个channel进行接收
	if cc.Config().Kafka.Enabled {
		var msgs []kafka.Message
		for _, e := range request.Events {
			data, err := msgpack.Marshal(e)
			if err != nil {
				return nil, err
			}
			// kafka保证了同个partition中是有序的
			msgs = append(msgs, kafka.Message{
				Key:   []byte(e.Honeypot),
				Value: data,
			})
		}
		err := gg.KafkaWriter.WriteMessages(ctx, msgs...)
		if err != nil {
			return &pb.Response{}, err
		}
	} else {
		for _, e := range request.Events {
			gg.EventChannel.AppendEvent(&model.HoneypotEvent{
				ID:        e.Id,
				Proto:     e.Proto,
				Honeypot:  e.Honeypot,
				Agent:     e.Agent,
				StartTime: time.Unix(e.StartTime.Seconds, 0),
				EndTime:   time.Unix(e.EndTime.Seconds, 0),
				SrcIp:     e.SrcIp,
				SrcPort:   int(e.SrcPort),
				DestIp:    e.DestIp,
				DestPort:  int(e.DestPort),
				RiskLevel: int(e.RiskLevel),
			})
		}
	}
	// 直接透传给judge节点
	if cc.Config().Judge.Enabled {
		var events []*model.HoneypotEvent
		for _, e := range request.Events {
			events = append(events, &model.HoneypotEvent{
				ID:        e.Id,
				Proto:     e.Proto,
				Honeypot:  e.Honeypot,
				Agent:     e.Agent,
				StartTime: time.Unix(e.StartTime.Seconds, 0),
				EndTime:   time.Unix(e.EndTime.Seconds, 0),
				SrcIp:     e.SrcIp,
				SrcPort:   int(e.SrcPort),
				DestIp:    e.DestIp,
				DestPort:  int(e.DestPort),
				RiskLevel: int(e.RiskLevel),
			})
		}
		go queue.Push2JudgeSendQueue(events)
	}
	return &pb.Response{}, nil
}

// GetEvent 获取事件信息
func (t *Event) GetEvent(ctx context.Context, request *pb.GetEventRequest) (*pb.GetEventResponse, error) {
	res, err := redis.Bytes(gg.RedisConnPool.Get().Do("GET", gg.HoneypotEventDetailRedisKey+request.EventId))
	if err != nil {
		log.Errorf("get event %s from redis failed %v", request.EventId, err)
		return &pb.GetEventResponse{}, errors.New(fmt.Sprintf("get event %s from redis failed %v", request.EventId, err))
	}
	var event model.HoneypotEvent
	err = msgpack.Unmarshal(res, &event)
	if err != nil {
		log.Errorf("unmarshal event %d failed %v", err)
		return &pb.GetEventResponse{}, errors.New(fmt.Sprintf("get event %s from redis failed %v", request.EventId, err))
	}
	event, err = db.GetHoneypotEvent(ctx, request.EventId)
	if err != nil {
		return &pb.GetEventResponse{}, errors.New(fmt.Sprintf("get event %s from ck failed %v", request.EventId, err))
	}
	return &pb.GetEventResponse{Event: &pb.HoneypotEvent{
		Id:        event.ID,
		Proto:     event.Proto,
		Honeypot:  event.Honeypot,
		Agent:     event.Agent,
		StartTime: &timestamppb.Timestamp{Seconds: event.StartTime.Unix()},
		EndTime:   &timestamppb.Timestamp{Seconds: event.EndTime.Unix()},
		SrcIp:     event.SrcIp,
		SrcPort:   int32(event.SrcPort),
		DestIp:    event.DestIp,
		DestPort:  int32(event.DestPort),
		RiskLevel: int32(event.RiskLevel),
	}}, nil
}

// ListEvent 获取事件列表
func (t *Event) ListEvent(ctx context.Context, request *pb.ListEventRequest) (*pb.ListEventResponse, error) {
	var resp pb.ListEventResponse
	events, err := db.QueryHoneypotEvents(ctx, db.QueryHoneypotEventParam{
		Limit:  request.Limit,
		Offset: request.Offset,
	})
	if err != nil {
		return &pb.ListEventResponse{}, errors.New(fmt.Sprintf("get event from ck failed %v", err))
	}
	resp.Events = make([]*pb.HoneypotEvent, 0)
	for _, event := range events {
		resp.Events = append(resp.Events, &pb.HoneypotEvent{
			Id:        event.ID,
			Proto:     event.Proto,
			Honeypot:  event.Honeypot,
			Agent:     event.Agent,
			StartTime: &timestamppb.Timestamp{Seconds: event.StartTime.Unix()},
			EndTime:   &timestamppb.Timestamp{Seconds: event.EndTime.Unix()},
			SrcIp:     event.SrcIp,
			SrcPort:   int32(event.SrcPort),
			DestIp:    event.DestIp,
			DestPort:  int32(event.DestPort),
			RiskLevel: int32(event.RiskLevel),
		})
	}
	return &resp, nil
}
