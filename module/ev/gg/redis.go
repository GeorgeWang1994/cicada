package gg

import (
	"github.com/GeorgeWang1994/cicada/module/ev/cc"
	"github.com/garyburd/redigo/redis"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

var RedisConnPool *redis.Pool

func InitRedisConnPool() {
	if !cc.Config().Redis.Enabled {
		return
	}

	auth, dsn := formatRedisAddr(cc.Config().Redis.Dsn)
	maxIdle := cc.Config().Redis.MaxIdle
	idleTimeout := 240 * time.Second

	connTimeout := time.Duration(cc.Config().Redis.ConnTimeout) * time.Millisecond
	readTimeout := time.Duration(cc.Config().Redis.ReadTimeout) * time.Millisecond
	writeTimeout := time.Duration(cc.Config().Redis.WriteTimeout) * time.Millisecond

	RedisConnPool = &redis.Pool{
		MaxIdle:     maxIdle,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(
				"tcp",
				dsn,
				redis.DialConnectTimeout(connTimeout),
				redis.DialReadTimeout(readTimeout),
				redis.DialWriteTimeout(writeTimeout),
			)
			if err != nil {
				return nil, err
			}
			if auth != "" {
				if _, err := c.Do("AUTH", auth); err != nil {
					_ = c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: PingRedis,
	}
}

func formatRedisAddr(addrConfig string) (string, string) {
	if redisAddr := strings.Split(addrConfig, "@"); len(redisAddr) == 1 {
		return "", redisAddr[0]
	} else {
		return strings.Join(redisAddr[0:len(redisAddr)-1], "@"), redisAddr[len(redisAddr)-1]
	}
}

func PingRedis(c redis.Conn, t time.Time) error {
	_, err := c.Do("ping")
	if err != nil {
		log.Println("[ERROR] ping redis fail", err)
	}
	return err
}
