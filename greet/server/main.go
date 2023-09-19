package main

import (
	pb "github.com/lincolnjpg/grpc_go_course/greet/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.GreetServiceServer
}

var addr = "0.0.0.0:50051"

func main() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Could not listen", err)
		return
	}

	log.Printf("Listenin on %v", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(listener); err != nil {
		log.Fatalln("Could not serve", err)
	}
}
