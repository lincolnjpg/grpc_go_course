package main

import (
	"io"
	"log"

	pb "github.com/lincolnjpg/grpc_go_course/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg function was invoked")

	var res int
	var count int

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Avg: float64(res) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving: %v\n", req)

		res += int(req.Num)
		count += 1
	}
}
