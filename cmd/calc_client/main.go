package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Shodske/calc-grpc/pkg/calculator"
	"google.golang.org/grpc"
)

func main() {
	// Read env variables for host and port, default to localhost:50051
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := calculator.NewCalculatorClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	args := os.Args[1:]
	if len(args) != 3 {
		log.Fatal("Incorrect number of arguments, usage: add x y")
	}

	var res *calculator.Result
	x, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		log.Fatal(err)
	}

	switch args[0] {
	case "add":
		res, err = c.Add(ctx, &calculator.Values{X: x, Y: y})
	case "sum":
		fallthrough
	case "evaluate":
		log.Fatalf("Unimplemented command: %s", args[0])
	default:
		log.Fatalf("Unknown command: %s", args[0])
	}

	if err != nil {
		log.Fatalf("Error while sending gRPC request: %v", err)
	}
	log.Printf("Result: %f", res.Value)
}
