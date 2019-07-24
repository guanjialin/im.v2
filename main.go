package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"im.v2/router"
)

func main() {
	engine := gin.Default()

	router.Register(engine)

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

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("服务启动失败:", err)
	}

	log.Println("server closed")
}
