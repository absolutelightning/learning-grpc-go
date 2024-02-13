package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/absolutelightning/learning-grpc-go/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Println("Greet function was invoked with input %v", in.FirstName)
	return &pb.GreetResponse{Result: fmt.Sprintf("Hello, %s", in.FirstName)}, nil
}
