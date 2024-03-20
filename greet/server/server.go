package main

import (
	"context"
	"fmt"
	pb "grpc/greet/proto"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.GreetServerServer
}

func (*server) Greet(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Print("The client canceled the request")
			return nil, status.Error(
				codes.DeadlineExceeded,
				"The client canceled the request",
			)
		}
		time.Sleep(1 * time.Second)
	}

	return &pb.GetResponse{Result: "Hello " + in.GetFirstName()}, nil
}

func main() {
	addr := "localhost:50069"
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("An error occur %v", err)
	}

	fmt.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()

	pb.RegisterGreetServerServer(s, &server{})

	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("An error occur %v", err)
	}

}
