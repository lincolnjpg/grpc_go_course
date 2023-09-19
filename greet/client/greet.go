package main

import (
	"context"
	"log"

	pb "github.com/lincolnjpg/grpc_go_course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "Junin"})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
		return
	}

	log.Printf("Greeting: %v\n", res.Result)
}
