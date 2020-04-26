package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-redis/redis/v7"
	"go.uber.org/zap"
)

const _redisKey = "VALUE"

func startHTTPServer() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)

}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	client := redis.NewClient(&redis.Options{
		Addr: "redis.service.consul:6379",
	})

	for {
		time.Sleep(5 * time.Second)

		val, err := client.Get(_redisKey).Result()
		if err != nil {
			logger.Error("failed to get key",
				zap.Error(err),
				zap.String("host", "redis"),
			)

			// initialize the key
			err = client.Set(_redisKey, 0, 0).Err()
			if err != nil {
				logger.Fatal("failed to iniialize key",
					zap.Error(err),
				)
			}

			continue
		}

		logger.Info("increment",
			zap.String("value", val),
		)

		err = client.Incr(_redisKey).Err()
		if err != nil {
			logger.Error("failed to increment key",
				zap.Error(err),
			)
		}
	}

}
