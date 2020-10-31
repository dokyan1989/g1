package main

import (
	"context"
	"log"

	"github.com/dokyan1989/g1/app/service1/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewProductServiceClient(conn)
	resp, err := client.ListProducts(context.Background(), &pb.ListProductsRequest{})
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}
	log.Println(resp)
}