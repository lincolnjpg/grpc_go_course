package main

import (
	"log"
	"net"

	pb "github.com/lincolnjpg/grpc_go_course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Server struct {
	pb.GreetServiceServer
}

var addr = "localhost:50051"

func main() {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Could not listen", err)
		return
	}

	log.Printf("Listenin on %v", addr)

	opts := []grpc.ServerOption{}
	tls := true

	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("Failed to server: %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(listener); err != nil {
		log.Fatalln("Could not serve", err)
	}
}
