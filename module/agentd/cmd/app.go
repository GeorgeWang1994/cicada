package main

import (
	"github.com/GeorgeWang1994/cicada/module/agentd/cc"
	"github.com/GeorgeWang1994/cicada/module/agentd/cron"
	"github.com/GeorgeWang1994/cicada/module/agentd/gg"
)

func initApp() error {
	err := cc.ParseConfig("")
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

	return nil
}
