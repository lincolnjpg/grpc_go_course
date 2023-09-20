package main

import (
	"log"
	"net"

	pb "github.com/lincolnjpg/grpc_go_course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.CalculatorServiceServer
}

var addr = "0.0.0.0:50051"

func main() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Could not listen: %v\n", err)
		return
	}

	log.Printf("Listenin on %v", addr)
	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})
	reflection.Register(s)

	if err = s.Serve(listener); err != nil {
		log.Fatalf("Could not serve: %v\n", err)
	}
}
