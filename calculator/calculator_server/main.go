package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/sushant102004/gRPC-Learning/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	calculatorpb.CalculatorServiceServer
}

func (*GRPCServer) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResult, error) {
	a := req.GetA()
	b := req.GetB()

	res := calculatorpb.SumResult{
		// Here SumResult is not a data type it is a key that object is having.
		SumResult: a + b,
	}

	fmt.Println("Sum Function Executed ~ Result: ", res.SumResult)

	return &res, nil
}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &GRPCServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
