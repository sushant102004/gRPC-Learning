package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/sushant102004/gRPC-Learning/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("0.0.0.0:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	// unary(c)
	doStream(c)
}

func unary(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Sushant",
			LastName:  "Dhiman",
		},
	}

	resp, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Result)

}

func doStream(c greetpb.GreetServiceClient) {
	req := &greetpb.GreetStreamRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Sushant",
		},
	}
	res, err := c.GreetStream(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(msg)
	}

}
