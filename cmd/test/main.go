package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Long-Software/Bex/packages/db"
	"github.com/Long-Software/Bex/packages/env"
	"github.com/Long-Software/Bex/packages/log"
)

type Config struct {
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
}

func main() {
	var c Config
	var logger log.ConsoleLogger
	err := env.LoadENVConfig(&c, ".env")
	if err != nil {
		logger.NewLog(log.FATAL, err.Error())
		os.Exit(0)
	}

	client, err := db.NewRedis(&db.RedisOptions{
		Addr:     fmt.Sprintf("%s:%s", c.RedisHost, c.RedisPort),
		Password: c.RedisPassword,
		DB:       0,
	})
	if err != nil {
		logger.NewLog(log.FATAL, err.Error())
		os.Exit(0)
	}

	err = client.Insert(context.Background(), "name", "long", time.Millisecond)
	if err != nil {
		logger.NewLog(log.ERROR, err.Error())
		os.Exit(0)
	}


	name, err := client.Read(context.Background(), "name")
	if err != nil {
		logger.NewLog(log.FATAL, err.Error())
		os.Exit(0)
	}
	fmt.Println(name)
}
