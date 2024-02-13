package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"

	pb "github.com/absolutelightning/learning-grpc-go/calculator/proto"
)

func main() {
	var serverAddr = "0.0.0.0:7777"

	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}

	c := pb.NewCalculatorServiceClient(conn)

	doAverage(c)

	defer conn.Close()
}

func doAverage(c pb.CalculatorServiceClient) {
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Failed to call Average: %v", err)
	}
	for i := 0; i < 5; i++ {
		fmt.Println("Enter number")
		var num int32
		fmt.Scanf("%d", &num)
		err := stream.Send(&pb.AverageRequest{Number: num})
		if err != nil {
			log.Fatalf("Failed to send number: %v", err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to receive response: %v", err)
	}
	fmt.Println("Average result =", res.Result)
}

func doSum(c pb.CalculatorServiceClient) {
	for {
		fmt.Println("Enter first number")
		var a, b int32
		fmt.Scanf("%d", &a)
		fmt.Println("Enter second number")
		fmt.Scanf("%d", &b)
		res, err := c.Sum(context.Background(), &pb.SumRequest{A: a, B: b})
		if err != nil {
			log.Fatalf("Failed to call Sum: %v", err)
		}
		fmt.Println("Sum result =", res.Result)
	}
}

func doPrimeFactor(c pb.CalculatorServiceClient) {
	for {
		var num int32
		fmt.Println("Enter number to prime factorize")
		fmt.Scanf("%d", &num)
		stream, err := c.PrimeFactors(context.Background(), &pb.PrimeFactorsRequest{Number: num})
		if err != nil {
			log.Fatalf("Failed to call Sum: %v", err)
		}
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				break
			}
			fmt.Println(res.Factors)
		}
	}

}
