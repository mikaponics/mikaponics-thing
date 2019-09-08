package controllers

import (
	"context"
	// "log"

	// tspb "github.com/golang/protobuf/ptypes/timestamp"

	pb "github.com/mikaponics/mikaponics-thing/api"
)

func (s *MikaponicsThingServer) SetThing(ctx context.Context, in *pb.SetThingRequest) (*pb.SetThingResponse, error) {
	return &pb.SetThingResponse{
		Message: "Thing was created",
		Status: true,
	}, nil
}
