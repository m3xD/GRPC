package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc/AVG/proto"
	"log"
	"time"
)

func getAVG(client pb.GetAVGClient) {

	streamRequest := []*pb.GetRequest{
		{Num: 1},
		{Num: 2},
		{Num: 3},
		{Num: 4},
	}

	stream, err := client.AGV(context.Background())

	if err != nil {
		log.Fatalln("Error", err)
	}

	for _, num := range streamRequest {
		fmt.Println("Sending request:", num)
		stream.Send(num)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalln("Error", err)
	}

	fmt.Printf("Ans:%f", res.Num)
}

func main() {
	cl, err := grpc.Dial("localhost:50069", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln("Error", err)
	}
	defer cl.Close()

	client := pb.NewGetAVGClient(cl)

	getAVG(client)
}
