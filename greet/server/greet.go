package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/absolutelightning/learning-grpc-go/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Println("Greet function was invoked with input %v", in.FirstName)
	return &pb.GreetResponse{Result: fmt.Sprintf("Hello, %s", in.FirstName)}, nil
}

func (s *Server) GreetManyTimes(in *pb.GreetRequest, streamp pb.GreetService_GreetManyTimesServer) error {
	log.Println("GreetManyTimes function was invoked with input %v", in.FirstName)
	for i := 0; i < 10; i++ {
		streamp.Send(&pb.GreetResponse{Result: fmt.Sprintf("Hello, %s %d", in.FirstName, i)})
	}
	return nil
}

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked")
	res := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{Result: res})
		}
		if err != nil {
			log.Fatal("error")
		}

		res += fmt.Sprintf("Hello, %s! \n", req.FirstName)
	}
}

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatal("error")
		}

		res := "Hello, " + req.FirstName + "!"
		err = stream.Send(&pb.GreetResponse{Result: res})

		if err != nil {
			log.Fatal("error")
		}
	}

}
