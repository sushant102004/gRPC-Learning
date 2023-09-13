package main

import (
	"context"
	"log"

	"github.com/sushant102004/gRPC-Learning/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial(":3000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	req := &calculatorpb.SumRequest{
		A: 4,
		B: 2,
	}

	c.Sum(context.Background(), req)

}
