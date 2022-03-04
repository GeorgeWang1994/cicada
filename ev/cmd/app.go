package main

import (
	"cicada/ev/cc"
	"cicada/ev/gg"
	"cicada/ev/rpc"
	"cicada/ev/sender"
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

	if cc.Config().Kafka.Enabled {
		gg.InitKafka()
	}

	if cc.Config().Clickhouse.Enable {
		gg.InitClickhouseClient(ctx)
	}

	if cc.Config().Judge.Enabled {
		sender.Start(ctx)
	}

	return nil
}
