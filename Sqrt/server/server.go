package main

import (
	"context"
	"fmt"
	Proto "grpc/Sqrt/proto"
	"log"
	"math"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	Proto.GetSQRTServer
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50069")

	if err != nil {
		log.Fatalf("An error has occured %v", err)
	}

	fmt.Print("Listening on addr localhost:50069")

	// create new server
	s := grpc.NewServer()

	Proto.RegisterGetSQRTServer(s, &server{})

	err = s.Serve(lis)

	if err != nil {
		log.Fatalln("Something went wrong", err)
	}

}

func (*server) SQRT(ctx context.Context, in *Proto.NumRequest) (*Proto.NumRespone, error) {

	num := in.GetNum()

	if num < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %d", num),
		)
	}

	return &Proto.NumRespone{Num: math.Sqrt(float64(num))}, nil
}
