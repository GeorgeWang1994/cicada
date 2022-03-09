package gg

import (
	"cicada/ev/cc"
	"cicada/ev/sender/queue"
	"cicada/ev/store/db"
	"cicada/pkg/model"
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack"
)

func InitWorker(ctx context.Context) {
	for i := 0; i < cc.Config().EventWorker.InitCap; i++ {
		go EventWorkerRun(ctx)
	}
}

func EventWorkerRun(ctx context.Context) {
	for {
		if cc.Config().Kafka.Enabled {
			msg, err := KafkaReader.ReadMessage(context.Background())
			if err != nil {
				log.Errorf("read kafka message failed %v", err)
			}

			var event *model.HoneypotEvent
			err = msgpack.Unmarshal(msg.Value, event)
			if err != nil {
				log.Errorf("unmarshal kafka message failed %v", err)
			}

			if cc.Config().Judge.Enabled {
				queue.Push2JudgeSendQueue([]*model.HoneypotEvent{event})
			}

			if cc.Config().Clickhouse.Enable {
				err := db.AsyncBatchInsertHoneypotEvent(ctx, []*model.HoneypotEvent{event}, false)
				if err != nil {
					log.Errorf("worker run insert failded %v", err)
				}
			}
		}
	}
}
