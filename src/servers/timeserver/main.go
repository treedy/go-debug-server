package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"
	pb "timeservices"

	"google.golang.org/grpc"
)

func getUTCTime() string {
	return fmt.Sprintf("UTC Time is %s", time.Now())
}

type server struct{}

func (s *server) GetUTCTime(ctx context.Context, in *pb.UTCTimeRequest) (*pb.UTCTimeResponse, error) {
	return &pb.UTCTimeResponse{Time: getUTCTime()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to lesten: %v", err)
	}

	svr := grpc.NewServer()
	pb.RegisterUTCTimeStringServer(svr, &server{})
	if err := svr.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
