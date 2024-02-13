package main

import (
	"context"
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
