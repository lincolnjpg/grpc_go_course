package main

import (
	pb "github.com/lincolnjpg/grpc_go_course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
		return
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)
	//doSum(c)
	//doPrimes(c)
	doAvg(c)
}
