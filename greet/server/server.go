package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	addr := "localhost:50000"
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("An error occur %v", err)
	}

	fmt.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()

	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("An error occur %v", err)
	}

}
