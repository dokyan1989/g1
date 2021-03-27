package transformer

import (
	"github.com/dokyan1989/g1/app/service1/internal/store"
	"github.com/dokyan1989/g1/app/service1/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

/*--------------------------------------
| Employee
--------------------------------------*/

func TransformEmployee2Proto(e store.Employee) *pb.Employee {
	return &pb.Employee{
		EmpNo:     e.EmpNo,
		BirthDate: timestamppb.New(e.BirthDate),
		FirstName: e.FirstName,
		LastName:  e.LastName,
		Gender:    TranformString2GenderEnum(e.Gender),
		HireDate:  timestamppb.New(e.HireDate),
	}
}

func TransformEmployeeList2Proto(el []store.Employee) []*pb.Employee {
	elPb := make([]*pb.Employee, len(el))

	for i, e := range el {
		elPb[i] = TransformEmployee2Proto(e)
	}

	return elPb
}

func TranformString2GenderEnum(gender string) pb.Gender {
	switch gender {
	case "M":
		return pb.Gender_MALE
	case "F":
		return pb.Gender_FEMALE
	}

	return pb.Gender_UNKNOWN
}
