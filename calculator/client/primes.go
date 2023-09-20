package main

import (
	"context"
	"io"
	"log"

	pb "github.com/lincolnjpg/grpc_go_course/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient) {
	log.Println("doGreetManyTimes was invoked")
	req := &pb.PrimesRequest{
		Num: 120,
	}
	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while Primes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
			return
		}

		log.Printf("Primes: %d\n", msg.Response)
	}
}
