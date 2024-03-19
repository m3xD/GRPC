package main

import (
	"google.golang.org/grpc"
	pb "grpc/AVG/proto"
	"io"
	"log"
	"net"
)

type server struct {
	pb.GetAVGServer
}

func (*server) AGV(str pb.GetAVG_AGVServer) error {
	var num float32 = 0
	var count float32 = 0
	for {
		msg, err := str.Recv()

		if err == io.EOF {
			str.SendAndClose(&pb.GetResponse{
				Num: num / count,
			})
			return nil
		}

		if err != nil {
			log.Fatalln("ror", err)
		}
		count++
		num += float32(msg.Num)
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50069")

	if err != nil {
		log.Fatalln("Error", err)
	}

	s := grpc.NewServer()

	pb.RegisterGetAVGServer(s, &server{})

	err = s.Serve(lis)

	if err != nil {
		log.Fatalln("Error", err)
	}
}
