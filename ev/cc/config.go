package cc

import (
	"cicada/pkg/utils/file"
	"encoding/json"
	"errors"
	"log"
	"sync"
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

type ClusterNode struct {
	Addrs []string `json:"addrs"`
}

type JudgeConfig struct {
	Enabled     bool                    `json:"enabled"`
	Batch       int                     `json:"batch"`
	ConnTimeout int                     `json:"connTimeout"`
	CallTimeout int                     `json:"callTimeout"`
	MaxConns    int                     `json:"maxConns"`
	MaxIdle     int                     `json:"maxIdle"`
	Replicas    int                     `json:"replicas"`
	Cluster     map[string]string       `json:"cluster"`
	ClusterList map[string]*ClusterNode `json:"clusterList"`
}

type GlobalConfig struct {
	Debug      bool              `json:"debug"`
	Rpc        *RpcConfig        `json:"rpc"`
	Judge      *JudgeConfig      `json:"judge"`
	Clickhouse *ClickhouseConfig `json:"clickhouse"`
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