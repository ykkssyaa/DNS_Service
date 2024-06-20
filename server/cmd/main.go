package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/ykkssyaa/DNS_Service/server/internal/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"net/http"
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gen.RegisterDnsServiceHandlerFromEndpoint(ctx, mux, "localhost:12201", opts)

	if err != nil {
		return err
	}

	return http.ListenAndServe(":8081", mux)
}

func main() {
	if err := run(); err != nil {
		grpclog.Fatal(err)
	}
}
