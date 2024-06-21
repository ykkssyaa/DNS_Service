package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/ykkssyaa/DNS_Service/server/gen"
	"github.com/ykkssyaa/DNS_Service/server/internal/server"
	"github.com/ykkssyaa/DNS_Service/server/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

func runRest(lg *logger.Logger) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := gen.RegisterDnsServiceHandlerFromEndpoint(ctx, mux, "localhost:12201", opts)

	if err != nil {
		lg.Err.Fatal(err)
	}
	lg.Info.Printf("rest server listening at 8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		lg.Err.Fatal(err)
	}
}

func runGrpc(lg *logger.Logger) {

	lis, err := net.Listen("tcp", ":12201")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	dnsServer := server.NewServer(lg)
	gen.RegisterDnsServiceServer(grpcServer, dnsServer)

	lg.Info.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		lg.Err.Fatalf("failed to serve: %v", err)
	}
}

func main() {

	lg := logger.InitLogger()

	go runRest(lg)
	runGrpc(lg)
}
