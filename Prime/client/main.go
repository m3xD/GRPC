package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc/Prime/proto"
	"io"
	"log"
)

func getPrime(client pb.GetPrimeClient) {
	fmt.Print("Has been invoked")
	rq := &pb.GetRequest{
		Num: 50,
	}

	stream, err := client.Prime(context.Background(), rq)

	if err != nil {
		log.Fatalln("An error occurred", err)
	}
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln("An error occurred", err)
		}
		fmt.Printf("%d\n", msg.Num)
	}
}

func main() {
	cli, err := grpc.Dial("localhost:50069", grpc.WithInsecure())

	if err != nil {
		log.Fatalln("Has an error", err)
	}

	defer cli.Close()

	actualClient := pb.NewGetPrimeClient(cli)

	getPrime(actualClient)
}
