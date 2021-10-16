package main

import (
	"log"

	"github.com/stalin-777/test-apiq/config"
	"github.com/stalin-777/test-apiq/logger"
	"github.com/stalin-777/test-apiq/server"
)

func main() {

	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	err = logger.InitZapLogger(
		cfg.Logger.Path,
		cfg.Logger.FileName,
		cfg.Logger.MaxSize,
		cfg.Logger.MaxBackups,
		cfg.Logger.MaxAge,
	)
	if err != nil {
		log.Fatal(err)
	}

	server.Run(cfg)
}
