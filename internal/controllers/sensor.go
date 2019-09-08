package controllers

import (
	"context"
	// "log"

	// tspb "github.com/golang/protobuf/ptypes/timestamp"

	pb "github.com/mikaponics/mikaponics-thing/api"
)

func (s *MikaponicsThingServer) SetSensor(ctx context.Context, in *pb.SetSensorRequest) (*pb.SetSensorResponse, error) {
	return &pb.SetSensorResponse{
		Message: "Sensor was created",
		Status: true,
	}, nil
}
