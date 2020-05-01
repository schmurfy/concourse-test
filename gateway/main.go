package main

import (
	"encoding/json"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
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

	conn, err := grpc.Dial(
		cfg.ServiceAddr,
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}
	serviceClient = pb.NewServiceClient(conn)
	defer conn.Close()

	http.HandleFunc("/addreses", func(w http.ResponseWriter, r *http.Request) {
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

	http.ListenAndServe(cfg.ListenAddr, nil)
}

// func HelloServer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
// }
