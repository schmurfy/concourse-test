package main

import (
	"encoding/json"
	"net/http"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/schmurfy/concourse-test/gateway/generated_pb/service"
)

var (
	serviceClient pb.ServiceClient
)

func main() {
	conn, err := grpc.Dial(
		"127.0.0.1:8000",
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}
	serviceClient = pb.NewServiceClient(conn)
	defer conn.Close()

	http.HandleFunc("/addreses", func(w http.ResponseWriter, r *http.Request) {
		ret, err := serviceClient.GetAddress(r.Context(), &empty.Empty{})
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

	http.ListenAndServe(":8080", nil)
}

// func HelloServer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
// }
