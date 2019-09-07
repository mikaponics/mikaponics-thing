package controllers

import (
	"context"
	// "log"

	// tspb "github.com/golang/protobuf/ptypes/timestamp"

	pb "github.com/mikaponics/mikaponics-thing/api"
)

func (s *MikaponicsRPC) SetTimeSeriesDatum(ctx context.Context, in *pb.SetTimeSeriesDatumRequest) (*pb.SetTimeSeriesDatumResponse, error) {
	return &pb.SetTimeSeriesDatumResponse{
		Message: "Instrument was created",
		Status: true,
	}, nil
}
