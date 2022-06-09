package main

import (
	"context"
	"flag"
	"github.com/GeorgeWang1994/cicada/module/judge/cc"
	"github.com/GeorgeWang1994/cicada/module/judge/cron"
	"github.com/GeorgeWang1994/cicada/module/judge/gg"
	"github.com/GeorgeWang1994/cicada/module/judge/rpc"
)

func initApp() error {
	cfg := flag.String("c", "cfg.json", "configuration file")
	err := cc.ParseConfig(*cfg)
	if err != nil {
		return err
	}

	ctx := context.Background()
	rpc.Start(ctx)

	if cc.Config().Alarm.Enabled && cc.Config().Alarm.Redis.Enabled {
		gg.InitRedisConnPool()
	}

	// 定期同步数据
	go cron.SyncAlarmStrategy()
	go cron.SyncSubscribeStrategy()

	return nil
}
