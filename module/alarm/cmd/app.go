package main

import (
	"flag"
	"github.com/GeorgeWang1994/cicada/module/alarm/cc"
	"github.com/GeorgeWang1994/cicada/module/alarm/gg"
	"github.com/GeorgeWang1994/cicada/module/alarm/msg"
)

func initApp() error {
	cfg := flag.String("c", "config.json", "configuration file")
	flag.Parse()
	err := cc.ParseConfig(*cfg)
	if err != nil {
		return err
	}

	gg.InitRedisConnPool()
	go msg.Consume()

	return nil
}
