package main

import (
	"context"
	pb "github.com/lincolnjpg/grpc_go_course/calculator/proto"
	"log"
)

func (s *Server) Calculate(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Calculate function was invoked with %v\n", in)
	return &pb.CalculatorResponse{
		Result: in.Num1 + in.Num2,
	}, nil
}
