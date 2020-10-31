package main

import (
	"log"
	"net"

	"github.com/dokyan1989/g1/app/tourgrpc/internal/pb"
	"github.com/dokyan1989/g1/app/tourgrpc/internal/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, server.NewHelloServer())
	grpcServer.Serve(lis)
}
