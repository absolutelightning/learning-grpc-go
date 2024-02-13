package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	pb "github.com/absolutelightning/learning-grpc-go/calculator/proto"
)

func main() {
	var serverAddr = "52.38.203.20:7777"

	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}

	c := pb.NewCalculatorServiceClient(conn)

	for {
		fmt.Println("Enter first number")
		var a, b int32
		fmt.Scanf("%d", &a)
		fmt.Println("Enter second number")
		fmt.Scanf("%d", &b)
		res, _ := c.Sum(context.Background(), &pb.SumRequest{A: a, B: b})
		fmt.Println("Sum result =", res.Result)
	}

	defer conn.Close()
}
