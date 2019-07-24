package config

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init()  {
	if gin.IsDebugging() {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:true,
		FullTimestamp:true,
	})
}