package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "event",
		Short: "event daemon",
		Run: func(cmd *cobra.Command, args []string) {
			err := initApp()
			if err != nil {
				log.Fatal(err)
			}
		},
	}
	return cmd
}

func main() {
	if err := newRootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}
