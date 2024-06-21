package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"github.com/ykkssyaa/DNS_Service/server/gen"
	"github.com/ykkssyaa/DNS_Service/server/internal/consts"
	"github.com/ykkssyaa/DNS_Service/server/internal/server"
	"github.com/ykkssyaa/DNS_Service/server/pkg/config"
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
	err := gen.RegisterDnsServiceHandlerFromEndpoint(ctx, mux, "localhost:"+viper.GetString("ports.grpc"), opts)

	if err != nil {
		lg.Err.Fatal(err)
	}

	port := viper.GetString("ports.rest")

	lg.Info.Printf("rest server listening at " + port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		lg.Err.Fatal(err)
	}
}

func runGrpc(lg *logger.Logger) {

	lis, err := net.Listen("tcp", ":"+viper.GetString("ports.grpc"))
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

	err := config.InitConfig(consts.ConfigFilePath)
	if err != nil {
		lg.Err.Fatal(err)
	}

	go runRest(lg)
	runGrpc(lg)
}
