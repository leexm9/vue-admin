package main

import (
	"context"
	"fmt"
	"go_web/configs"
	"go_web/interaction/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	serverConf := configs.GetInstance().ServerConf

	server := &http.Server{
		Addr:    fmt.Sprintf("localhost:%d", serverConf.Port),
		Handler: router.Init(),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalln("Init Server failure, error:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-quit

	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln("Server shutdown error:", err)
	}
	log.Println("Server exiting")
}
