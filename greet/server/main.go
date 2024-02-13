package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/absolutelightning/learning-grpc-go/greet/proto"
)

var addr = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Failed to listen on address: ", addr)
	}
	log.Println("Listening on address: ", addr)

	s := grpc.NewServer()

	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: ", err)
	}
}
