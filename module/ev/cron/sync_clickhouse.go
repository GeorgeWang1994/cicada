package cron

import (
	"context"
	"github.com/GeorgeWang1994/cicada/module/ev/cc"
	"github.com/GeorgeWang1994/cicada/module/ev/gg"
	"github.com/GeorgeWang1994/cicada/module/ev/store/db"
	"github.com/GeorgeWang1994/cicada/module/pkg/model"
	"github.com/garyburd/redigo/redis"
	log "github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack"
	"time"
)

func Sync2Clickhouse(ctx context.Context) {
	for {
		if cc.Config().Clickhouse.Enable {
			// todo: 确认没有处理成功的需要把数据塞回Redis
			endTime := time.Now().Unix() - 60*10
			// 从过去的10分钟开始算，开始更新之前30分钟的数据
			startTime := endTime - 60*30
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
					err := db.BatchInsertHoneypotEvent(ctx, events)
					if err != nil {
						log.Errorf("worker run insert failded %v", err)
					} else {
						// 成功后删除redis中的数据
						_, err := gg.RedisConnPool.Get().Do("ZREMRANGE", gg.HoneypotEventRedisKey, startTime, endTime)
						if err != nil {
							log.Errorf("get redis events failed %v", err)
						}
					}
				}
			}
		}
	}
}
