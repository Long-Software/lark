package main

import (
	"fmt"
	"os"

	"github.com/Long-Software/Bex/cmd/shorten-url/internal/api"
	"github.com/Long-Software/Bex/cmd/shorten-url/internal/config"
	"github.com/Long-Software/Bex/packages/db"
	"github.com/Long-Software/Bex/packages/env"
	"github.com/Long-Software/Bex/packages/log"
)

func main() {
	var cfg config.Config
	var cl log.ConsoleLogger

	err := env.LoadENVConfig(cfg, ".env")
	if err != nil {
		cl.NewLog(log.FATAL, err.Error())
		os.Exit(0)
	}
	store, err := db.NewRedis(&db.RedisOptions{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       0,
	})
	defer store.Close()
	server, err := api.NewServer(cfg, *store)
	if err != nil {
		cl.NewLog(log.FATAL, err.Error())
		os.Exit(0)
	}

	err = server.Start(fmt.Sprintf(":%d", cfg.AppPort))
	if err != nil {
		cl.NewLog(log.FATAL, err.Error())
		os.Exit(0)
	}
}
