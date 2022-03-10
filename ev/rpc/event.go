package rpc

import (
	"cicada/ev/cc"
	"cicada/ev/gg"
	"cicada/ev/store/db"
	"cicada/pkg/model"
	pb "cicada/proto/api/ev"
	"context"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/rogpeppe/fastuuid"
	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

var kafkaUuidGen = fastuuid.MustNewGenerator()

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
	if cc.Config().Kafka.Enabled {
		var msgs []kafka.Message
		for _, e := range request.Events {
			uid := kafkaUuidGen.Hex128()
			data, err := msgpack.Marshal(e)
			if err != nil {
				return nil, err
			}
			msgs = append(msgs, kafka.Message{
				Key:   []byte(uid),
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
				ID:         e.Id,
				Proto:      e.Proto,
				Honeypot:   e.Honeypot,
				Agent:      e.Agent,
				StartTime:  time.Unix(e.StartTime.Seconds, 0),
				EndTime:    time.Unix(e.EndTime.Seconds, 0),
				SrcIp:      e.SrcIp,
				SrcPort:    int(e.SrcPort),
				SrcMac:     e.SrcMac,
				DestIp:     e.DestIp,
				DestPort:   int(e.DestPort),
				EventTypes: e.EventTypes,
				RiskLevel:  int(e.RiskLevel),
			})
		}
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
		Id:         event.ID,
		Proto:      event.Proto,
		Honeypot:   event.Honeypot,
		Agent:      event.Agent,
		StartTime:  &timestamppb.Timestamp{Seconds: event.StartTime.Unix()},
		EndTime:    &timestamppb.Timestamp{Seconds: event.EndTime.Unix()},
		SrcIp:      event.SrcIp,
		SrcPort:    int32(event.SrcPort),
		SrcMac:     event.SrcMac,
		DestIp:     event.DestIp,
		DestPort:   int32(event.DestPort),
		EventTypes: event.EventTypes,
		RiskLevel:  int32(event.RiskLevel),
	}}, nil
}

// ListEvent 获取事件列表
func (t *Event) ListEvent(ctx context.Context, request *pb.ListEventRequest) (*pb.ListEventResponse, error) {
	var resp *pb.ListEventResponse
	return resp, nil
}
