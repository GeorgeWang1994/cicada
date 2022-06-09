package main

import (
	"context"
	"flag"
	"github.com/GeorgeWang1994/cicada/ev/cc"
	"github.com/GeorgeWang1994/cicada/ev/cron"
	"github.com/GeorgeWang1994/cicada/ev/gg"
	"github.com/GeorgeWang1994/cicada/ev/rpc"
	"github.com/GeorgeWang1994/cicada/ev/sender"
)

func initApp() error {
	cfg := flag.String("c", "cfg.json", "configuration file")
	err := cc.ParseConfig(*cfg)
	if err != nil {
		return err
	}

	ctx := context.Background()
	rpc.Start(ctx)

	if cc.Config().Redis.Enabled {
		gg.InitRedisConnPool()
	}

	if cc.Config().Kafka.Enabled {
		gg.InitKafka()
	}

	if cc.Config().Clickhouse.Enable {
		gg.InitClickhouseClient(ctx)
		cron.Sync2Clickhouse(ctx)
	}

	if cc.Config().EventWorker.Enable {
		gg.InitWorker(ctx)
	}

	if cc.Config().Judge.Enabled {
		sender.Start(ctx)
	}

	return nil
}
