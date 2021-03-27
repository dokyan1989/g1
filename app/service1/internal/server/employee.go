package server

import (
	"context"

	"github.com/dokyan1989/g1/app/service1/internal/store"
	"github.com/dokyan1989/g1/app/service1/internal/transformer"
	"github.com/dokyan1989/g1/app/service1/pb"
	"github.com/dokyan1989/g1/lib/utils"
)

func (s *Server) ListEmployees(ctx context.Context, in *pb.ListEmployeesRequest) (*pb.ListEmployeesResponse, error) {
	employees, err := s.store.ListEmployees(ctx, store.ListEmployeesParams{
		EmpNos:        in.EmpNos,
		FromBirthDate: utils.ConvertTimestampPb2Time(in.FromBirthDate),
		ToBirthDate:   utils.ConvertTimestampPb2Time(in.ToBirthDate),
		Names:         in.Names,
		Gender:        transformer.TranformGenderEnum2String(in.Gender),
		FromHireDate:  utils.ConvertTimestampPb2Time(in.FromHireDate),
		ToHireDate:    utils.ConvertTimestampPb2Time(in.ToHireDate),
		Limit:         in.Limit,
		Offset:        in.Offset,
	})
	if err != nil {
		return nil, err
	}

	return &pb.ListEmployeesResponse{
		Data: transformer.TransformEmployeeList2Proto(employees),
	}, nil
}

// CreateEmployee ...
func (s *Server) CreateEmployee(ctx context.Context, in *pb.CreateEmployeeRequest) (*pb.CreateEmployeeResponse, error) {
	lastId, err := s.store.CreateEmployee(ctx, store.CreateEmployeeParams{
		BirthDate: utils.ConvertTimestampPb2Time(in.BirthDate),
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Gender:    transformer.TranformGenderEnum2String(in.Gender),
		HireDate:  utils.ConvertTimestampPb2Time(in.HireDate),
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateEmployeeResponse{
		EmpNo: lastId,
	}, nil
}

func (s *Server) UpdateEmployee(ctx context.Context, in *pb.UpdateEmployeeRequest) (*pb.UpdateEmployeeResponse, error) {
	lastId, err := s.store.UpdateEmployee(ctx, store.UpdateEmployeeParams{
		EmpNo:     in.EmpNo,
		BirthDate: utils.ConvertTimestampPb2Time(in.BirthDate),
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Gender:    transformer.TranformGenderEnum2String(in.Gender),
		HireDate:  utils.ConvertTimestampPb2Time(in.HireDate),
	})
	if err != nil {
		return nil, err
	}

	return &pb.UpdateEmployeeResponse{
		EmpNo: lastId,
	}, nil
}

func (s *Server) DeleteEmployee(ctx context.Context, in *pb.DeleteEmployeeRequest) (*pb.DeleteEmployeeResponse, error) {
	lastId, err := s.store.DeleteEmployee(ctx, in.EmpNo)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteEmployeeResponse{
		EmpNo: lastId,
	}, nil
}
