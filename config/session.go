package config

import "sync"

type sessionConfig struct {
	SecretKey string `json:"secret_key"`
	Idle      int    `json:"idle"`
	once      sync.Once
}

var session sessionConfig

func Session() *sessionConfig {
	redis.once.Do(func() {
		parseFromJson(config[fileSession], &session)
	})

	return &session
}
