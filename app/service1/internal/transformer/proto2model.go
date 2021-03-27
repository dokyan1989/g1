package transformer

import "github.com/dokyan1989/g1/app/service1/pb"

func TranformGenderEnum2String(gender pb.Gender) string {
	switch gender {
	case pb.Gender_MALE:
		return "M"
	case pb.Gender_FEMALE:
		return "F"
	}

	return ""
}
