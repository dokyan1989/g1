package server

import (
	"context"

	"github.com/dokyan1989/g1/app/service1/pb"
)

func (s *Server) ListEmployees(context.Context, *pb.ListEmployeesRequest) (*pb.ListEmployeesResponse, error) {
	return &pb.ListEmployeesResponse{}, nil
}
