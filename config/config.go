package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

var configDir string
var configOnce sync.Once

const (
	filePg = iota
	fileRedis
	fileSession
	fileGithub
)

var config = map[int]string{
	filePg:      "/pg.json",
	fileRedis:   "/redis.json",
	fileSession: "/session.json",
	fileGithub:  "/github.json",
}

func init() {
	configOnce.Do(func() {
		exist := false
		if configDir, exist = os.LookupEnv("config-im"); !exist {
			panic("请设置配置文件环境变量 config-im")
		}

		for _, file := range config {
			if _, err := os.Stat(configDir + file); os.IsNotExist(err) {
				panic(fmt.Sprintf("配置文件 %s 不存在", file))
			}
		}
	})
}

func parseFromJson(file string, model interface{}) {
	content, err := ioutil.ReadFile(configDir + file)
	if err != nil {
		panic(fmt.Sprintf("读取配置文件[%s]失败: %s", file, err.Error()))
	}

	if err := json.Unmarshal(content, &model); err != nil {
		panic(fmt.Sprintf("解析json失败:%s", err.Error()))
	}
}
