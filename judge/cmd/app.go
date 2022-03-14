package main

import (
	"cicada/judge/cc"
	"cicada/judge/gg"
	"cicada/judge/rpc"
	"context"
	"flag"
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

	return nil
}
