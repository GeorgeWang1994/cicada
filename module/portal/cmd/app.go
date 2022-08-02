package main

import (
	"context"
	"flag"
	"github.com/GeorgeWang1994/cicada/module/portal/cc"
	"github.com/GeorgeWang1994/cicada/module/portal/gg"
	"github.com/GeorgeWang1994/cicada/module/portal/rpc"
)

func initApp() error {
	cfg := flag.String("c", "config.json", "configuration file")
	flag.Parse()
	err := cc.ParseConfig(*cfg)
	if err != nil {
		return err
	}

	ctx := context.Background()
	rpc.Start(ctx)

	if cc.Config().Redis.Enabled {
		gg.InitRedisConnPool()
	}

	// 定期同步数据
	//go cron.SyncAlarmStrategy()
	//go cron.SyncSubscribeStrategy()

	return nil
}
