package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/stalin-777/test-apiq/internal/config"
	"github.com/stalin-777/test-apiq/internal/logger"
	"github.com/stalin-777/test-apiq/internal/router"
	"github.com/stalin-777/test-apiq/internal/tasker"
)

func main() {

	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	err = logger.InitZapLogger(cfg.Logger)
	if err != nil {
		log.Fatal(err)
	}

	server, err := router.New(cfg)
	if err != nil {
		logger.Fatalf("Failed to ger router, error:%s", err.Error())
	}

	tasker := tasker.New(cfg.WorkersNum)
	tasker.Run(server, cfg.WorkersNum)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go server.Run(cfg)

	<-quit

	if err := server.Shutdown(); err != nil {
		logger.Fatal(err)
	}

	tasker.CancelFunc()
	tasker.WG.Wait()

	logger.Info("---------Server Stopped---------")
}
