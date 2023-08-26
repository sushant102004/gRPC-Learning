package main

import (
	"log"
	"net"

	"github.com/sushant102004/gRPC-Learning/greet/greetpb"
	"google.golang.org/grpc"
)

type Server struct {
	greetpb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Error setting up TCP listner: ", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error serving server: ", err)
	}
}
