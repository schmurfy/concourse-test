package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/blendle/zapdriver"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/schmurfy/concourse-test/gateway/config"
	pb "github.com/schmurfy/concourse-test/gateway/generated_pb/service"
)

var (
	serviceClient pb.ServiceClient
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

	conn, err := grpc.Dial(
		cfg.ServiceAddr,
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}
	serviceClient = pb.NewServiceClient(conn)
	defer conn.Close()

	http.HandleFunc("/addrs", func(w http.ResponseWriter, r *http.Request) {
		ret, err := serviceClient.GetAddresses(r.Context(), &empty.Empty{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(ret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(data)
	})

	http.HandleFunc("/redis", func(w http.ResponseWriter, r *http.Request) {
		_, err := serviceClient.IncrementRedis(r.Context(), &pb.IncrementRedisRequest{
			Key: "test",
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	})

	logger.Info("gateway started",
		zap.String("address", cfg.ListenAddr),
		zap.String("service", cfg.ServiceAddr),
	)

	http.ListenAndServe(cfg.ListenAddr, nil)
}

// func HelloServer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
// }
