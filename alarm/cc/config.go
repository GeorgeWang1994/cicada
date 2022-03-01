package cc

import (
	"cicada/pkg/utils/file"
	"encoding/json"
	"errors"
	"log"
	"sync"
)

type WorkerConfig struct {
	IM   int `json:"im"`
	Sms  int `json:"sms"`
	Mail int `json:"mail"`
}

type RpcConfig struct {
	RpcAddr string `json:"rpc_addr"`
	Timeout int    `json:"timeout"`
}

type RedisConfig struct {
	Dsn          string `json:"dsn"`
	MaxIdle      int    `json:"maxIdle"`
	ConnTimeout  int    `json:"connTimeout"`
	ReadTimeout  int    `json:"readTimeout"`
	WriteTimeout int    `json:"writeTimeout"`
	// 不同威胁等级的队列
	HighQueues   []string `json:"highQueues"`
	MediumQueues []string `json:"mediumQueues"`
	LowQueues    []string `json:"lowQueues"`
	// 不用providor的队列
	UserIMQueue   string `json:"userIMQueue"`
	UserSmsQueue  string `json:"userSmsQueue"`
	UserMailQueue string `json:"userMailQueue"`
}

type GlobalConfig struct {
	Debug bool         `json:"debug"`
	RPC   *RpcConfig   `json:"rpc"`
	Redis *RedisConfig `json:"redis"`
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
