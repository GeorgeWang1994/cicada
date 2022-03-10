package cron

import (
	"cicada/ev/cc"
	"cicada/ev/gg"
	"cicada/ev/store/db"
	"cicada/pkg/model"
	"context"
	"github.com/garyburd/redigo/redis"
	log "github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack"
	"time"
)

func Sync2Clickhouse(ctx context.Context) {
	for {
		if cc.Config().Clickhouse.Enable {
			endTime := time.Now().Unix()
			startTime := endTime - 60*10
			res, err := redis.Strings(gg.RedisConnPool.Get().Do("ZRANGE", gg.HoneypotEventRedisKey, startTime, endTime))
			if err != nil {
				log.Errorf("get redis events failed %v", err)
			} else {
				if len(res) == 0 {
					log.Error("empty events from redis")
					continue
				}
				var events []*model.HoneypotEvent
				for _, d := range res {
					var event *model.HoneypotEvent
					err = msgpack.Unmarshal([]byte(d), &event)
					if err != nil {
						log.Errorf("unmarshal kafka message failed %v", err)
						continue
					}
					events = append(events, event)
				}
				if len(events) > 0 {
					err := db.AsyncBatchInsertHoneypotEvent(ctx, events, false)
					if err != nil {
						log.Errorf("worker run insert failded %v", err)
					}
				}
			}
		}
	}
}
