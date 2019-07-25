package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"im.v2/config"
	"sync"
)

var redisConn redis.Conn
var redisOnce sync.Once

func Redis() redis.Conn {
	redisOnce.Do(func() {
		var err error
		redisConn, err = redis.Dial("tcp", config.Redis().Addr,
			redis.DialPassword(config.Redis().Password))
		if err != nil {
			logrus.Panic("连接redis失败:", err)
			return
		}
	})

	return redisConn
}