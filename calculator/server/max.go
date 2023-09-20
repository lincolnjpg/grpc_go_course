package main

import (
	"io"
	"log"

	pb "github.com/lincolnjpg/grpc_go_course/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max function was invoked")

	var max int64

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		if req.Num > max {
			max = req.Num
			err = stream.Send(&pb.MaxResponse{
				Response: max,
			})
		}

		if err != nil {
			log.Fatalf("Error while sending data to client: %v\n", err)
		}
	}
}
