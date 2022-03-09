package rpc

import (
	"cicada/ev/cc"
	"cicada/ev/gg"
	"cicada/pkg/model"
	pb "cicada/proto/api"
	"context"
	"errors"
	"github.com/rogpeppe/fastuuid"
	"github.com/segmentio/kafka-go"
	"github.com/vmihailenco/msgpack"
	"time"
)

var kafkaUuidGen = fastuuid.MustNewGenerator()

type Event struct{}

func (t *Event) Ping(ctx context.Context, resp *model.RpcResponse) error {
	return nil
}

// ReceiveEvent 接收来自探针的事件
func (t *Event) ReceiveEvent(ctx context.Context, request *pb.EventRequest) (*pb.EventResponse, error) {
	if len(request.Events) <= 0 {
		return &pb.EventResponse{}, errors.New("empty events")
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
			return &pb.EventResponse{}, err
		}
	} else {
		for _, e := range request.Events {
			gg.EventWorker.AppendEvent(&model.HoneypotEvent{
				ID:         e.Id,
				Proto:      e.Proto,
				Honeypot:   e.Honeypot,
				Agent:      e.Agent,
				StartTime:  time.Unix(int64(e.StartTime), 0),
				EndTime:    time.Unix(int64(e.EndTime), 0),
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
	return &pb.EventResponse{}, nil
}
