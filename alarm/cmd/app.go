package main

import (
	"cicada/alarm/cc"
	"cicada/alarm/gg"
	"cicada/alarm/msg"
	"flag"
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
