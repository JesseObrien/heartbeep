package main

import (
	"context"
	"time"

	"github.com/apex/log"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"google.golang.org/grpc"

	"github.com/jesseobrien/heartbeep/internal/beeps"
)

func main() {
	log.Info("beeper online")

	address := "localhost:8888"

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect to collector: %v", err)
	}

	log.Info("connected to beep collector")

	defer conn.Close()

	cl := beeps.NewBeepCollectorClient(conn)

	for range time.Tick(time.Second) {
		ctx := context.Background()
		request := &beeps.BeepRequest{Time: ptypes.TimestampNow(), RequestId: uuid.New().String()}
		resp, err := cl.Beep(ctx, request, &grpc.EmptyCallOption{})
		if err != nil {
			log.WithError(err).Info("could not process beep response")
		}
		log.WithField("request_id", request.GetRequestId()).Info("Beeper beeped")
		responseTime, _ := ptypes.Timestamp(resp.GetTime())
		log.WithFields(log.Fields{"request_id": resp.RequestId, "time": responseTime}).Infof("received request back")
	}

}
