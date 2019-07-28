package config

import (
	"sync"
)

type redisConfig struct {
	Addr     string    `json:"addr"` // host:port
	Password string    `json:"password"`
	Database string    `json:"database"`
	once     sync.Once
}

var redis redisConfig

func Redis() *redisConfig {
	redis.once.Do(func() {
		parse(config[fileRedis], &redis)
	})

	return &redis
}