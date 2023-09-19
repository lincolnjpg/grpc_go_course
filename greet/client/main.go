package main

import (
	"log"

	"google.golang.org/grpc"
)

var addr = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr)
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}
	defer conn.Close()
}
