package cc

import (
	"cicada/pkg/utils/file"
	"encoding/json"
	"errors"
	"log"
	"sync"
	"time"
)

type RpcConfig struct {
	RpcAddr string `json:"rpc_addr"`
	Timeout int    `json:"timeout"`
}

type ClickhouseConfig struct {
	Enable      bool   `json:"enable"`
	Addr        string `json:"addr"`
	Database    string `json:"database"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	ConnTimeout int    `json:"connTimeout"`
}

type JudgeConfig struct {
	Enabled      bool              `json:"enabled"`
	Batch        int               `json:"batch"`
	DialTimeout  int               `json:"dialTimeout"`
	IdleTimeout  int               `json:"idleTimeout"`
	ReadTimeout  int               `json:"readTimeout"`
	WriteTimeout int               `json:"writeTimeout"`
	InitCap      int               `json:"initCap"`
	MaxCap       int               `json:"MaxCap"`
	Replicas     int               `json:"replicas"`
	Cluster      map[string]string `json:"cluster"`
}

type KafkaConfig struct {
	Enabled    bool          `json:"enabled"`
	Broker     []string      `json:"broker"`
	Topic      string        `json:"topic"`
	Partition  int           `json:"partition"`
	BatchSize  uint          `json:"batchSize"`
	Timeout    uint          `json:"timeout"`
	BatchDelay time.Duration `json:"batchDelay"`
	Compress   bool          `json:"compress"`
}

type EventWorkerConfig struct {
	Enable  bool `json:"enable"`
	DataCap int  `json:"dataCap"`
	InitCap int  `json:"initCap"`
	MaxCap  int  `json:"maxCap"`
}

type GlobalConfig struct {
	Debug       bool               `json:"debug"`
	Rpc         *RpcConfig         `json:"rpc"`
	Judge       *JudgeConfig       `json:"judge"`
	Kafka       *KafkaConfig       `json:"kafka"`
	Clickhouse  *ClickhouseConfig  `json:"clickhouse"`
	EventWorker *EventWorkerConfig `json:"eventWorker"`
}

var (
	lock sync.RWMutex
	cc   *GlobalConfig
)

func ParseConfig(cfg string) error {
	if cfg == "" {
		return errors.New("use -c to specify pivas file")
	}

	if !file.IsExist(cfg) {
		return errors.New("is not existent. maybe you need `mv cfg.example.json cfg.json`")
	}

	configContent, err := file.FileContent(cfg)
	if err != nil {
		return err
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		return err
	}

	lock.Lock()
	defer lock.Unlock()

	cc = &c

	log.Println("read cc file:", cfg, "successfully")
	return nil
}

func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return cc
}
