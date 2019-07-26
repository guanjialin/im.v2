package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"im.v2/model"
	"im.v2/router"
)

func main() {
	engine := gin.Default()
	engine.Use()

	logrus.Infoln("init router...")
	router.Register(engine)

	logrus.Infoln("init database table...")
	if err := model.InitTable(); err != nil {
		logrus.Panic("init database table failure:", err)
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

	logrus.Infoln("server start at :8000")
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("server start failure:", err)
	}

	logrus.Infoln("server stop...")
}
