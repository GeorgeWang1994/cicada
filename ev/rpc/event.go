package rpc

import (
	"cicada/ev/cc"
	"cicada/ev/store/db"
	"cicada/pkg/model"
	"context"
)

type Event int

func (t *Event) Ping(req model.NullRpcRequest, resp *model.RpcResponse) error {
	return nil
}

func (t *Event) Receive(context context.Context, args []*model.HoneypotEvent, reply *model.RpcResponse) error {

	if cc.Config().Judge.Enabled {
		//queue.Push2JudgeSendQueue(args)
	}

	if cc.Config().Clickhouse.Enable {
		err := db.AsyncBatchInsertHoneypotEvent(context, args, false)
		if err != nil {
			return err
		}
	}

	return nil
}
