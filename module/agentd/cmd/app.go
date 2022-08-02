package main

import (
	"flag"
	"github.com/GeorgeWang1994/cicada/module/agentd/cc"
	"github.com/GeorgeWang1994/cicada/module/agentd/cron"
	"github.com/GeorgeWang1994/cicada/module/agentd/gg"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func initApp() error {
	cfg := flag.String("c", "config.json", "configuration file")
	flag.Parse()
	err := cc.ParseConfig(*cfg)
	if err != nil {
		return err
	}

	worker := cc.Config().Worker
	if worker == 0 {
		worker = gg.DefaultWorker
	}

	for i := 0; i < worker; i++ {
		go cron.SendEvent()
	}

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		select {
		case sig := <-c:
			{
				log.Infof("Got %s signal. Aborting...\n", sig)
				os.Exit(1)
			}
		}
	}()

	return nil
}
