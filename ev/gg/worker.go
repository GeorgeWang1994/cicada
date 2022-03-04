package gg

import (
	"cicada/ev/cc"
	"cicada/ev/sender/queue"
	"cicada/ev/store/db"
	"cicada/pkg/model"
	"context"
	"github.com/sirupsen/logrus"
)

type eventChannel struct {
	eventChannel chan *model.HoneypotEvent
	closeChannel chan bool
}

var (
	EventWorker *eventChannel
)

func InitWorker(ctx context.Context) {
	EventWorker = &eventChannel{
		eventChannel: make(chan *model.HoneypotEvent, cc.Config().EventWorker.DataCap),
	}
	for i := 0; i < cc.Config().EventWorker.InitCap; i++ {
		go EventWorkerRun(ctx)
	}
}

func EventWorkerRun(ctx context.Context) {
	for {
		select {
		case data := <-EventWorker.eventChannel:
			// todo: 这样是否还有必要发到队列中，确认确认下实际的性能
			if cc.Config().Judge.Enabled {
				queue.Push2JudgeSendQueue(data)
			}

			if cc.Config().Clickhouse.Enable {
				err := db.AsyncBatchInsertHoneypotEvent(ctx, args, false)
				if err != nil {
					logrus.Errorf("worker run insert failded %v", err)
				}
			}
		case ctx.Done():

			return
		}
	}
}
