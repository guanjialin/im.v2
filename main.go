package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"im.v2/model"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"im.v2/router"
)

func main() {
	engine := gin.Default()
	engine.Use()

	logrus.Infoln("开始设置路由...")
	router.Register(engine)

	logrus.Infoln("开始创建数据库表...")
	if err := model.CreateTable(); err != nil {
		logrus.Panic("创建数据库表失败:", err)
		return
	}

	server := http.Server{
		Addr:    ":8000",
		Handler: engine,
	}

	go func() {
		sigInt := make(chan os.Signal, 1)
		signal.Notify(sigInt, os.Interrupt)
		<-sigInt

		if err := server.Shutdown(context.Background()); err != nil {
			log.Println("server shutdown:", err)
		}
	}()

	logrus.Infoln("开始启动服务...")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("服务启动失败:", err)
	}

	log.Println("server closed")
}
