package main

import (
	"context"
	"log"
	"time"

	pb "github.com/lincolnjpg/grpc_go_course/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doLongGreet was invoked")
	reqs := []*pb.AvgRequest{
		{
			Num: 7,
		},
		{
			Num: 12,
		},
		{
			Num: 3,
		},
	}
	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Avg: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Avg: %v\n", err)
	}

	log.Printf("Avg: %f\n", res.Avg)
}
