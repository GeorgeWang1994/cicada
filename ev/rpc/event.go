package rpc

import (
	"cicada/ev/cc"
	"cicada/ev/gg"
	"cicada/pkg/model"
	"context"
	"github.com/rogpeppe/fastuuid"
	"github.com/segmentio/kafka-go"
)

var kafkaUuidGen = fastuuid.MustNewGenerator()

type Event struct{}

func (t *Event) Ping(req model.NullRpcRequest, resp *model.RpcResponse) error {
	return nil
}

// ReceiveEvent 接收来自探针的事件
func (t *Event) ReceiveEvent(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	if cc.Config().Kafka.Enabled {
		uid := kafkaUuidGen.Hex128()
		err := gg.KafkaWriter.WriteMessages(ctx, kafka.Message{
			Key:   []byte(uid),
			Value: []byte(""),
		})
		if err != nil {
			return &pb.Response{}, nil
		}
	} else {

	}
	return &pb.Response{}, nil
}
