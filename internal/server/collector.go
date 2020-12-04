package server

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/apex/log"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"

	pb "github.com/jesseobrien/heartbeep/internal/beeps"
)

type CollectorServer struct {
	pb.UnimplementedBeepCollectorServer
}

func (bs *CollectorServer) Beep(ctx context.Context, beep *pb.BeepRequest) (*pb.BeepResponse, error) {
	now := time.Now()
	timeSent, _ := ptypes.Timestamp(beep.Time)
	log.WithField("time diff", now.Sub(timeSent)).Infof("Collector received beep")

	return &pb.BeepResponse{
		Time:      beep.GetTime(),
		RequestId: beep.GetRequestId(),
	}, nil
}

func (bs *CollectorServer) Run(port string) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBeepCollectorServer(s, &CollectorServer{})

	log.Info("beeper server starting...")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("error server could not start: %v", err)
	}
}
