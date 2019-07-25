package session

import (
	"github.com/boj/redistore"
	"github.com/sirupsen/logrus"
	"im.v2/config"
)

var store *redistore.RediStore

func Session() *redistore.RediStore {
	var err error
	store, err = redistore.NewRediStore(config.Session().Idle, "tcp",
		config.Redis().Addr, config.Redis().Password, []byte(config.Session().SecretKey))
	if err != nil {
		logrus.Panic("初始化session失败:", err)
		return nil
	}

	return store
}