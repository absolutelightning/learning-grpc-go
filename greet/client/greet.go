package main

import (
	"context"
	"io"
	"log"
	"time"

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

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("Do long greet was invoked")
	reqs := []*pb.GreetRequest{{
		FirstName: "absl",
	}, {
		FirstName: "absl1",
	}}

	stream, err := c.LongGreet(context.Background())

	for _, req := range reqs {
		if err != nil {
			log.Fatal("error")
		}
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("error")
	}
	log.Println("Result = %v", res.Result)
}

func doGreetEveryOne(c pb.GreetServiceClient) {
	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatal("error")
	}
	reqs := []*pb.GreetRequest{{FirstName: "absl"}, {FirstName: "absl1"}}

	waitCh := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Println("Sending message %v", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal("error")
				break
			}
			log.Println("Result = %v", res.Result)
		}
		close(waitCh)
	}()
	<-waitCh
}
