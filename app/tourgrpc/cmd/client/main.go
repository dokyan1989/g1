package main

import (
	"context"
	"log"

	"github.com/dokyan1989/g1/app/tourgrpc/internal/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	reply, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "An"})
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}

	log.Println(reply)
}
