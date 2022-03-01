package main

import (
	"cicada/ev/cc"
	"cicada/ev/gg"
	"cicada/ev/rpc"
	"flag"
)

func initApp() error {
	cfg := flag.String("c", "cfg.json", "configuration file")
	err := cc.ParseConfig(*cfg)
	if err != nil {
		return err
	}

	go rpc.Start()

	if cc.Config().Clickhouse.Enable {
		gg.InitClickhouseClient()
	}

	return nil
}
