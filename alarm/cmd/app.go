package main

import (
	"cicada/alarm/cc"
	"cicada/alarm/gg"
	"flag"
)

func initApp() error {
	cfg := flag.String("c", "cfg.json", "configuration file")
	err := cc.ParseConfig(*cfg)
	if err != nil {
		return err
	}

	gg.InitRedisConnPool()

	return nil
}
