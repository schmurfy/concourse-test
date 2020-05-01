package main

import (
	"log"

	"github.com/blendle/zapdriver"
	"github.com/schmurfy/concourse-test/service/server"
	"go.uber.org/zap"
)

func main() {
	logger, err := zapdriver.NewDevelopment(zap.AddStacktrace(zap.ErrorLevel))
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	server.Start(logger)
}
