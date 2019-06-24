package main

import (
	"log"
	"net"

	"github.com/Shodske/calc-grpc/pkg/calculator"
	"github.com/Shodske/calc-grpc/pkg/server"
	"google.golang.org/grpc"
)

func main() {
	log.Print("starting calculator gRPC server")

	lis, err := net.Listen("tcp", ":50051")
	defer lis.Close()
	if err != nil {
		log.Fatalf("failed to open tcp connection: %v", err)
	}

	s := grpc.NewServer()
	calculator.RegisterCalculatorServer(s, server.NewServer())
	defer s.Stop()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
