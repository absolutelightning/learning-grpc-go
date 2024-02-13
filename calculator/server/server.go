package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
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
