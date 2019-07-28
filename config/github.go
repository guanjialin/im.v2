package config

import (
	"sync"
)

type githubConfig struct {
	AccessTokenURL string `json:"access_token_url"`
	ClientID       string `json:"client_id"`
	ClientSecret   string `json:"client_secret"`
	once           sync.Once
}

var github githubConfig

func Github() *githubConfig {
	github.once.Do(func() {
		parse(config[fileGithub], &github)
	})

	return &github
}
