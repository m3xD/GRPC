package main

import (
	"fmt"
	"google.golang.org/grpc"
	Proto "grpc/Prime/proto"
	"log"
	"net"
)

type server struct {
	Proto.GetPrimeServer
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50069")

	if err != nil {
		log.Fatalf("An error has occured %v", err)
	}

	fmt.Print("Listening on addr localhost:50069")

	// create new server
	s := grpc.NewServer()

	Proto.RegisterGetPrimeServer(s, &server{})

	err = s.Serve(lis)

	if err != nil {
		log.Fatalln("Something went wrong", err)
	}

}

func (s *server) Prime(rq *Proto.GetRequest, srv Proto.GetPrime_PrimeServer) error {
	fmt.Print("Invoked")
	var curNumber = int(rq.GetNum())
	for i := 2; i < curNumber; i++ {
		for curNumber%i == 0 {
			curNumber /= i
			err := srv.Send(&Proto.GetResponse{
				Num: int32(i),
			})
			if err != nil {
				log.Fatalf("An error %v", err)
			}
		}
	}
	return nil
}
