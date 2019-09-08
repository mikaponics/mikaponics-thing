package controllers

import (
	"context"
	// "log"

	// tspb "github.com/golang/protobuf/ptypes/timestamp"

	pb "github.com/mikaponics/mikaponics-thing/api"
)

func (s *MikaponicsRPC) SetTimeSeriesDatum(ctx context.Context, in *pb.SetTimeSeriesDatumRequest) (*pb.SetTimeSeriesDatumResponse, error) {
	return &pb.SetTimeSeriesDatumResponse{
		Message: "Datum was created",
		Status: true,
	}, nil
}

func (s *MikaponicsRPC) SetTimeSeriesData(ctx context.Context, in *pb.SetTimeSeriesDataRequest) (*pb.SetTimeSeriesDataResponse, error) {
	return &pb.SetTimeSeriesDataResponse{
		Message: "Data was created",
		Status: true,
	}, nil
}
