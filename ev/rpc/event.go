package rpc

import (
	"cicada/ev/cc"
	"cicada/ev/sender/queue"
	"cicada/ev/store/db"
	"cicada/pkg/model"
	"context"
)

type Event struct{}

func (t *Event) Ping(req model.NullRpcRequest, resp *model.RpcResponse) error {
	return nil
}

// ReceiveEvent 接收来自探针的事件
func (t *Event) ReceiveEvent(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	if cc.Config().Judge.Enabled {
		queue.Push2JudgeSendQueue(args)
	}

	if cc.Config().Clickhouse.Enable {
		err := db.AsyncBatchInsertHoneypotEvent(ctx, args, false)
		if err != nil {
			return &pb.Response{}, nil
		}
	}

	return &pb.Response{}, nil
}
