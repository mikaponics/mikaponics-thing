package controllers

import (
	"context"
	"log"

	// tspb "github.com/golang/protobuf/ptypes/timestamp"

	pb "github.com/mikaponics/mikaponics-thing/api"
)


func (s *MikaponicsRPC) SetTimeSeriesData(ctx context.Context, in *pb.SetTimeSeriesDataRequest) (*pb.SetTimeSeriesDataResponse, error) {

	for _, v := range in.Data {
		log.Printf("%v", v)
	}

	return &pb.SetTimeSeriesDataResponse{
		Message: "Data was created",
		Status: true,
	}, nil
}
