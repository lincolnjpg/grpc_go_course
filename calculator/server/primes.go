package main

import (
	pb "github.com/lincolnjpg/grpc_go_course/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error {
	numerator := in.Num
	denominator := uint64(2)

	for {
		if numerator <= 1 {
			break
		}

		if numerator%denominator == 0 {
			stream.Send(&pb.PrimesResponse{
				Response: denominator,
			})
			numerator /= denominator
		} else {
			denominator += 1
		}
	}

	return nil
}
