package main

import (
	"context"
	"log"

	pb "github.com/lincolnjpg/grpc_go_course/calculator/proto"
)

func doCalculate(c pb.CalculatorServiceClient) {
	log.Println("doCalculate was invoked")

	res, err := c.Calculate(context.Background(), &pb.CalculatorRequest{
		Num1: 8,
		Num2: 12,
	})
	if err != nil {
		log.Fatalf("Could not calculate: %v\n", err)
		return
	}

	log.Printf("Calculate's response: %v\n", res.Result)
}
