package main

import (
	"context"
	"flag"
	"github.com/GeorgeWang1994/cicada/module/ev/cc"
	"github.com/GeorgeWang1994/cicada/module/ev/cron"
	"github.com/GeorgeWang1994/cicada/module/ev/gg"
	"github.com/GeorgeWang1994/cicada/module/ev/rpc"
	"github.com/GeorgeWang1994/cicada/module/ev/sender"
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

	if cc.Config().Kafka.Enabled {
		gg.InitKafka()
	}

	if cc.Config().Clickhouse.Enabled {
		gg.InitClickhouseClient(ctx)
		cron.Sync2Clickhouse(ctx)
	}

	if cc.Config().EventWorker.Enabled {
		gg.InitWorker(ctx)
	}

	if cc.Config().Judge.Enabled {
		sender.Start(ctx)
	}

	return nil
}
