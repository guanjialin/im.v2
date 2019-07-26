package config

import "sync"

type githubConfig struct {
	URL          string `json:"url"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	once         sync.Once
}

var github githubConfig

func Github() *githubConfig {
	github.once.Do(func() {
		parseFromJson(config[fileGithub], &github)
	})

	return &github
}
