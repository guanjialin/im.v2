package config

import "sync"

type pgConfig struct {
	Addr     string    `json:"addr"` // host:port
	User     string    `json:"user"`
	Password string    `json:"password"`
	Database string    `json:"database"`
	once     sync.Once
}

var pg pgConfig

func PG() *pgConfig {
	pg.once.Do(func() {
		parseFromJson(config[filePg], &pg)
	})

	return &pg
}