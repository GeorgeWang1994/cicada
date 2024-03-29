package cc

import (
	"encoding/json"
	"errors"
	"github.com/GeorgeWang1994/cicada/module/pkg/utils/file"
	log "github.com/sirupsen/logrus"
	"sync"
)

type RedisConfig struct {
	Enabled      bool   `json:"enabled"`
	Dsn          string `json:"dsn"`
	MaxIdle      int    `json:"maxIdle"`
	ConnTimeout  int    `json:"connTimeout"`
	ReadTimeout  int    `json:"readTimeout"`
	WriteTimeout int    `json:"writeTimeout"`
}

type RpcConfig struct {
	RpcAddr string `json:"rpc_addr"`
	Timeout int    `json:"timeout"`
}

type AlarmConfig struct {
	Enabled      bool         `json:"enabled"`
	MinInterval  int64        `json:"minInterval"`  // 告警最小间隔时间
	QueuePattern string       `json:"queuePattern"` // 告警的队列，用来告诉发往哪个redis key
	Redis        *RedisConfig `json:"redis"`
}

type PortalConfig struct {
	Servers      []string `json:"servers"`
	Timeout      int      `json:"timeout"`
	SyncInterval int64    `json:"sync_interval"` // 同步数据时间范围
}

type GlobalConfig struct {
	Rpc    *RpcConfig    `json:"rpc"`
	Alarm  *AlarmConfig  `json:"alarm"`
	Portal *PortalConfig `json:"portal"`
	Debug  bool          `json:"debug"`
}

var (
	lock sync.RWMutex
	cc   *GlobalConfig
)

func ParseConfig(cfg string) error {
	if cfg == "" {
		return errors.New("use -c to specify judge file")
	}

	if !file.IsExist(cfg) {
		return errors.New("is not existent. maybe you need `mv cfg.example.json config.json`")
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
