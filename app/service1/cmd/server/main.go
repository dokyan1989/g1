package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/dokyan1989/g1/app/service1/config"
	"github.com/dokyan1989/g1/app/service1/internal/server"
	"github.com/dokyan1989/g1/app/service1/pb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var cfg *config.Config

func main() {
	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	var err error
	stop := make(chan os.Signal, 1)
	cfg, err = config.Load()
	if err != nil {
		return err
	}

	grpcAddr := fmt.Sprintf("%s:%d", cfg.Server.GRPC.Host, cfg.Server.GRPC.Port)
	httpAddr := fmt.Sprintf("%s:%d", cfg.Server.HTTP.Host, cfg.Server.HTTP.Port)

	// run grpc server
	go func() {
		lis, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		grpcServer := grpc.NewServer()
		s := server.New(cfg)
		pb.RegisterEmployeeServiceServer(grpcServer, s)
		grpcServer.Serve(lis)
	}()

	// run grpc gateway
	go func() {
		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithInsecure()}
		if err := pb.RegisterEmployeeServiceHandlerFromEndpoint(context.Background(), mux, grpcAddr, opts); err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		http.ListenAndServe(httpAddr, mux)
	}()

	for {
		<-stop
		log.Println("Shutting down server")
	}
}
