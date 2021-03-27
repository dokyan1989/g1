package utils

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertTimestampPb2Time(tpb *timestamppb.Timestamp) time.Time {
	if tpb != nil {
		return tpb.AsTime()
	}

	return time.Time{}
}
