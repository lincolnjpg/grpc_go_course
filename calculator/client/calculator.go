package main

import (
	"context"
	"log"

	pb "github.com/lincolnjpg/grpc_go_course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  8,
		SecondNumber: 12,
	})
	if err != nil {
		log.Fatalf("Could not Sum: %v\n", err)
		return
	}

	log.Printf("Sum's response: %v\n", res.Result)
}
