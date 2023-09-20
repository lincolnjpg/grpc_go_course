package main

import (
	"context"
	"log"
	"time"

	pb "github.com/lincolnjpg/grpc_go_course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked with: %v\n", in)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			msg := "The client canceled the request"
			log.Println(msg)
			return nil, status.Error(codes.Canceled, msg)
		}

		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello, " + in.FirstName,
	}, nil
}
