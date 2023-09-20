package main

import (
	"context"
	"log"
	"time"

	pb "github.com/lincolnjpg/grpc_go_course/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doLongGreet was invoked")

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Avg: %v\n", err)
	}

	numbers := []int64{7, 12, 4}

	for _, number := range numbers {
		log.Printf("Sending req: %v\n", number)
		stream.Send(&pb.AvgRequest{
			Num: number,
		})
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Avg: %v\n", err)
	}

	log.Printf("Avg: %f\n", res.Avg)
}
