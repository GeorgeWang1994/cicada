package cc

import (
	"cicada/pkg/utils/file"
	"encoding/json"
	"errors"
	"log"
	"sync"
)

type GlobalConfig struct {
	ID       string `json:"id"`
	Debug    bool   `json:"debug"`
	RpcAddr  string `json:"rpc_addr"`
	Timeout  int    `json:"timeout"`
	Worker   int    `json:"worker"`    // 每次worker的数量
	PerCount int    `json:"per_count"` // 每次请求的个数
	Interval int    `json:"interval"`  // 间隔时间
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
