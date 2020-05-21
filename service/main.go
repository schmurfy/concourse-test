package main

import (
	"log"

	"github.com/blendle/zapdriver"
	"github.com/go-redis/redis/v7"
	"go.uber.org/zap"

	_ "github.com/schmurfy/test-lib2"

	"github.com/schmurfy/concourse-test/service/config"
	"github.com/schmurfy/concourse-test/service/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	logger, err := zapdriver.NewDevelopment(zap.AddStacktrace(zap.ErrorLevel))
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	client := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})

	logger.Info("service started",
		zap.String("address", cfg.ListenAddr),
		zap.String("redis", cfg.RedisAddr),
	)

	server.Start(logger, client)
}
