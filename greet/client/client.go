package main

import (
	"context"
	pb "grpc/Greet/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func getName(client pb.GreetServerClient, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	defer cancel()

	res, err := client.Greet(ctx, &pb.GetRequest{
		FirstName: "Khanh",
	})

	if err != nil {
		_, ok := status.FromError(err)

		if ok {
			log.Fatalln("An exxceed deadline")
			return
		} else {
			log.Fatalln("Not an grpc error")
		}
	}

	log.Println("The result is", res.Result)
}

func main() {
	c, err := grpc.Dial("localhost:50069", grpc.WithInsecure())

	if err != nil {
		log.Fatalln("Error", err)
	}

	defer c.Close()

	client := pb.NewGreetServerClient(c)

	getName(client, 3*time.Second)
}
