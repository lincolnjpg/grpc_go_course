package main

import (
	"log"

	pb "github.com/lincolnjpg/grpc_go_course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)
	//doGreet(c)
	//doGreetManyTimes(c)
	//doLongGreet(c)
	doGreetEveryone(c)
}
