package main

import (
	"context"
	"flag"
	"github.com/GeorgeWang1994/cicada/portal/cc"
	"github.com/GeorgeWang1994/cicada/portal/rpc"
)

func initApp() error {
	cfg := flag.String("c", "cfg.json", "configuration file")
	err := cc.ParseConfig(*cfg)
	if err != nil {
		return err
	}

	ctx := context.Background()
	rpc.Start(ctx)

	// 定期同步数据
	//go cron.SyncAlarmStrategy()
	//go cron.SyncSubscribeStrategy()

	return nil
}
