package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// 判断文件是否存在
// 如果存在返回true，不存在则返回false
func FileExist(path string) bool {
	_, err := os.Open(path)
	if err != nil {
		return os.IsExist(err)
	}

	return true
}

// 将配置文件path解析为struct类型的model
func FileParseToJson(path string, model interface{}) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(content, &model); err != nil {
		return err
	}

	return nil
}
