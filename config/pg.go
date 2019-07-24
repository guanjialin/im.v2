package config

import "sync"

type pg struct {
	Addr     string    `json:"addr"` // host:port
	User     string    `json:"user"`
	Password string    `json:"password"`
	Datbase string    `json:"database"`
	once     sync.Once
}

var pgConfig pg

func PG() *pg {
	pgConfig.once.Do(func() {
		parseFromJson(config[filePg], &pgConfig)
	})

	return &pgConfig
}