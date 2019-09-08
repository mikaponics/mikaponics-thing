package controllers

import (
	"context"
	// "log"

	// tspb "github.com/golang/protobuf/ptypes/timestamp"

	pb "github.com/mikaponics/mikaponics-thing/api"
)


func (s *MikaponicsThingServer) SetTimeSeriesData(ctx context.Context, in *pb.SetTimeSeriesDataRequest) (*pb.SetTimeSeriesDataResponse, error) {
    // // Iterate through all the data that was submitted by the embedded device.
	// for _, v := range in.Data {
	// 	// s.DAL.CreateTimeSeriesDatum(v.TenantId, v.SensorId, v.Value, v.Timestamp.Seconds)
	// 	log.Printf("SAVED: %v", v) // For debugging purposes only.
	// }

	return &pb.SetTimeSeriesDataResponse{
		Message: "Data was successfully uploaded",
		Status: true,
	}, nil
}
