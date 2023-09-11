package main

import (
	"context"
	"log"
	"net"

	"github.com/sushant102004/gRPC-Learning/greet/greetpb"
	"google.golang.org/grpc"
)

type Server struct {
	greetpb.GreetServiceServer
}

func (*Server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fName := req.GetGreeting().GetFirstName()
	lName := req.GetGreeting().GetLastName()
	result := "Hello, " + fName + " " + lName

	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
