package main

import (
	"fmt"
	"log"

	"github.com/sushant102004/gRPC-Learning/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	fmt.Printf("Client: %v", c)
}
