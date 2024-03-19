package main

import (
	pb "grpc/Max/proto"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.GetMaxServer
}

func (*server) Max(stream pb.GetMax_MaxServer) error {
	var cur int = 0
	for {
		str, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error %v", err)
		}
		if cur < int(str.Num) {
			cur = int(str.Num)
			err = stream.Send(&pb.ResponseNumber{
				Num: int32(cur),
			})
			if err != nil {
				log.Fatalf("Error %v\n", err)
			}
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50069")

	if err != nil {
		log.Fatalf("Error %v\n", err)
	}

	s := grpc.NewServer()

	pb.RegisterGetMaxServer(s, &server{})

	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("Error %v", err)
	}
}
