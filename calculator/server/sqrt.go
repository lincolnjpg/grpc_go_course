package main

import (
	"context"
	"fmt"
	"log"
	"math"

	pb "github.com/lincolnjpg/grpc_go_course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt was invoked with %v\n", in)

	number := in.Num

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %d\n", number),
		)
	}

	return &pb.SqrtResponse{
		Response: math.Sqrt(float64(number)),
	}, nil
}
