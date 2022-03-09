package gg

import (
	"cicada/ev/cc"
	"github.com/segmentio/kafka-go"
	"net"
	"time"
)

var KafkaWriter *kafka.Writer
var KafkaReader *kafka.Reader

// InitKafka 初始化kakfa客户端
func InitKafka() {
	transport := &kafka.Transport{
		Dial: (&net.Dialer{
			Timeout: time.Duration(cc.Config().Kafka.Timeout) * time.Second,
		}).DialContext,
	}

	KafkaWriter = &kafka.Writer{
		Addr:         kafka.TCP(cc.Config().Kafka.Broker...),
		Async:        true,
		Balancer:     &kafka.LeastBytes{},
		BatchSize:    int(cc.Config().Kafka.BatchSize),
		BatchTimeout: cc.Config().Kafka.BatchDelay,
		Topic:        cc.Config().Kafka.Topic,
		Transport:    transport,
	}

	KafkaReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:   cc.Config().Kafka.Broker,
		Topic:     cc.Config().Kafka.Topic,
		Partition: cc.Config().Kafka.Partition,
	})
}
