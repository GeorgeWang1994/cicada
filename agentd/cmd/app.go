package main

import (
	"cicada/agentd/cc"
	"cicada/agentd/cron"
	"cicada/agentd/gg"
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
