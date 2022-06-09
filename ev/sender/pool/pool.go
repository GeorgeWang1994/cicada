package pool

import (
	"github.com/GeorgeWang1994/cicada/ev/cc"
	"github.com/GeorgeWang1994/cicada/pkg/utils/pool"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"time"
)

var (
	JudgeConnPools *pool.GRPCPool
)

// InitConnPools 初始化各个组件的连接池
func InitConnPools() {
	cfg := cc.Config()

	// judge
	var targets []string
	for _, instance := range cfg.Judge.Cluster {
		targets = append(targets, instance)
	}
	options := &pool.Options{
		InitTargets:  targets,
		InitCap:      cfg.Judge.InitCap,
		MaxCap:       cfg.Judge.MaxCap,
		DialTimeout:  time.Duration(cfg.Judge.DialTimeout),
		IdleTimeout:  time.Duration(cfg.Judge.IdleTimeout),
		ReadTimeout:  time.Duration(cfg.Judge.ReadTimeout),
		WriteTimeout: time.Duration(cfg.Judge.WriteTimeout),
	}
	var err error
	JudgeConnPools, err = pool.NewGRPCPool(options, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
}
