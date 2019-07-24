package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var configOnce sync.Once
var configDir string

func init() {
	set := false
	configDir, set = os.LookupEnv("config-im")
	if !set {
		log.Fatalln("请设置配置文件环境变量config-im")
	}
}

func parseConfig(path string, mod interface{}) error {
	f, err := os.Open(configDir + path)
	if os.IsNotExist(err) {
		return errors.New("配置文件不存在")
	}

	config, err := ioutil.ReadAll(f)
	if err != nil {
		return errors.New("解析配置文件失败:" + err.Error())
	}

	err = json.Unmarshal(config, mod)
	if err != nil {
		return errors.New("解析JSON失败:" + err.Error())
	}

	return nil
}
