package gg

import (
	"context"
	"github.com/GeorgeWang1994/cicada/module/ev/cc"
	"github.com/GeorgeWang1994/cicada/module/pkg/model"
	log "github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack"
	"time"
)

const HoneypotEventRedisKey = "honeypot_events"
const HoneypotEventDetailRedisKey = "event_detail"

type (
	eventChannel struct {
		event chan *model.HoneypotEvent
		close chan bool
	}
)

var EventChannel *eventChannel

func (e *eventChannel) AppendEvent(event *model.HoneypotEvent) {
	e.event <- event
}

func (e *eventChannel) Close() {
	close(e.close)
}

func InitWorker(ctx context.Context) {
	for i := 0; i < cc.Config().EventWorker.InitCap; i++ {
		go EventWorkerRun(ctx)
	}
}

func EventWorkerRun(ctx context.Context) {
	log.Println("worker start...")
	for {
		if cc.Config().Kafka.Enabled {
			msg, err := KafkaReader.ReadMessage(ctx)
			if err != nil {
				log.Errorf("read kafka message failed %v", err)
			}
			var event model.HoneypotEvent
			err = msgpack.Unmarshal(msg.Value, &event)
			if err != nil {
				log.Errorf("unmarshal kafka message failed %v", err)
				continue
			}
			_, err = RedisConnPool.Get().Do("SET", HoneypotEventDetailRedisKey+event.ID, msg)
			if err != nil {
				log.Errorf("add kafka message to redis failed %v", err)
			}
			_, err = RedisConnPool.Get().Do("ZADD", HoneypotEventRedisKey, time.Now().Unix(), event.ID)
			if err != nil {
				log.Errorf("add kafka message to redis range failed %v", err)
			}
		} else {
			select {
			case event := <-EventChannel.event:
				data, err := msgpack.Marshal(event)
				if err != nil {
					log.Errorf("marshal event from channel failed %v", err)
				}
				_, err = RedisConnPool.Get().Do("SET", HoneypotEventDetailRedisKey+event.ID, data)
				if err != nil {
					log.Errorf("add kafka message to redis failed %v", err)
				}
				_, err = RedisConnPool.Get().Do("ZADD", HoneypotEventRedisKey, time.Now().Unix(), event.ID)
				if err != nil {
					log.Errorf("add kafka message to redis range failed %v", err)
				}
			case <-EventChannel.close:
				log.Println("channel close")
				return
			case <-ctx.Done():
				log.Println("context done")
				return
			}
		}
	}
}
