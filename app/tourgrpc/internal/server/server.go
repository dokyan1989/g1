package server

import (
	"context"
	"fmt"

	"github.com/dokyan1989/g1/app/tourgrpc/internal/pb"
)

// HelloServer ...
type HelloServer struct {
	pb.UnimplementedGreeterServer
}

// NewHelloServer ...
func NewHelloServer() *HelloServer {
	return &HelloServer{}
}

// SayHello ...
func (hs *HelloServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: fmt.Sprintf("Hello, %s", in.Name),
	}, nil
}
