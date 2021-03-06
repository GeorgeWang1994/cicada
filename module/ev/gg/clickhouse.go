package gg

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/compress"
	"github.com/GeorgeWang1994/cicada/module/ev/cc"
	log "github.com/sirupsen/logrus"
	"time"
)

var (
	ClickhouseClient clickhouse.Conn
)

// InitClickhouseClient 初始化CK客户端
func InitClickhouseClient(ctx context.Context) {
	var err error
	client, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{cc.Config().Clickhouse.Addr},
		Auth: clickhouse.Auth{
			Database: cc.Config().Clickhouse.Database,
			Username: cc.Config().Clickhouse.Username,
			Password: cc.Config().Clickhouse.Password,
		},
		ConnMaxLifetime: time.Hour,
		Compression:     &clickhouse.Compression{Method: compress.LZ4},
		DialTimeout:     time.Duration(cc.Config().Clickhouse.ConnTimeout),
	})
	if err != nil {
		log.Fatal(err)
	}
	ckCtx := clickhouse.Context(ctx, clickhouse.WithSettings(clickhouse.Settings{
		"max_block_size": 10,
	}), clickhouse.WithProgress(func(p *clickhouse.Progress) {
		fmt.Println("progress: ", p)
	}), clickhouse.WithProfileInfo(func(p *clickhouse.ProfileInfo) {
		fmt.Println("profile info: ", p)
	}))
	if err := client.Ping(ckCtx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			log.Fatalf("Catch Clickhouse exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
	}
	ClickhouseClient = client
}
