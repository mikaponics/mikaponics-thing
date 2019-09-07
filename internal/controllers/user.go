package controllers

import (
	"context"
	// "log"

	// tspb "github.com/golang/protobuf/ptypes/timestamp"

	pb "github.com/mikaponics/mikaponics-thing/api"
)

func (s *MikaponicsRPC) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{
		Key: "Instrument was updated",
		Status: true,
	}, nil
}
