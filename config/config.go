package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/sirupsen/logrus"

	"im.v2/utils"
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
			if !utils.FileExist(configDir + file) {
				panic(fmt.Sprintf("配置文件 %s 不存在", configDir+file))
			}
		}
	})
}

func parse(file string, model interface{}) {
	err := utils.FileParseToJson(configDir+file, model)
	if err != nil {
		logrus.Panicf("解析配置文件: %s 失败: %s", file, err)
		return
	}
	logrus.Debugf("config: %#v", model)
}