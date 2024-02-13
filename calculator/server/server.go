package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"

	pb "github.com/absolutelightning/learning-grpc-go/calculator/proto"
)

type Server struct {
	pb.CalculatorServiceServer
}

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Println("Add function was invoked with input %v", in)
	return &pb.SumResponse{Result: in.A + in.B}, nil
}

func (s *Server) PrimeFactors(req *pb.PrimeFactorsRequest, stream pb.CalculatorService_PrimeFactorsServer) error {
	log.Println("PrimeFactors function was invoked with input %v", req)
	num := req.Number
	k := int32(2)
	for num > 1 {
		if num%k == 0 {
			stream.Send(&pb.PrimeFactorsResponse{Factors: k})
			num = num / k
		} else {
			k = k + 1
		}
	}
	return nil
}

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max function was invoked")
	var max int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error when reading client stream: %v", err)
		}
		if req.Number > max {
			max = req.Number
			stream.SendMsg(&pb.MaxResponse{Result: max})
		}
	}
	return nil
}

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Println("Average function was invoked")
	var sum int32
	var count int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AverageResponse{Result: float64(sum*1.0) / float64(count)})
		}
		if err != nil {
			return err
		}
		sum += req.Number
		count++
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 7777))
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	err = s.Serve(lis)

	if err != nil {
		log.Fatal("failed to serve: ", err)
	}

}
