package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	pb "github.com/absolutelightning/learning-grpc-go/greet/proto"
)

var serverAddr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)

	defer conn.Close()
}
