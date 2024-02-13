package main

import (
	"context"
	"io"
	"log"

	pb "github.com/absolutelightning/learning-grpc-go/greet/proto"
)

func doGreet(client pb.GreetServiceClient) {
	log.Print("Do greet was invoked")
	res, err := client.Greet(context.Background(), &pb.GreetRequest{FirstName: "absl"})

	if err != nil {
		log.Fatal("error")
	}

	log.Println("Result = %v", res)
}

func doGreetManyTimes(client pb.GreetServiceClient) {
	log.Print("Do greet many times was invoked")

	req := pb.GreetRequest{FirstName: "absl"}
	stream, err := client.GreetManyTimes(context.Background(), &req)

	if err != nil {
		log.Fatal("error")
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("error")
		}

		log.Println("Result = %v", msg.Result)
	}
}
