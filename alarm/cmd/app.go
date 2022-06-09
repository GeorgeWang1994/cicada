package main

import (
	"flag"
	"github.com/GeorgeWang1994/cicada/alarm/cc"
	"github.com/GeorgeWang1994/cicada/alarm/gg"
	"github.com/GeorgeWang1994/cicada/alarm/msg"
)

func initApp() error {
	cfg := flag.String("c", "cfg.json", "configuration file")
	err := cc.ParseConfig(*cfg)
	if err != nil {
		return err
	}

	gg.InitRedisConnPool()
	go msg.Consume()

	return nil
}
