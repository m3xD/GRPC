package main

import (
	"context"
	pb "grpc/Sqrt/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func getSQRT(c pb.GetSQRTClient) {
	rc, err := c.SQRT(context.Background(), &pb.NumRequest{Num: -2})

	if err != nil {
		c, ok := status.FromError(err)

		if ok {
			log.Printf("A message from server %s\n", c.Message())
			log.Printf("A code from server %s\n", c.Code())
			if c.Code() == codes.InvalidArgument {
				log.Printf("We probably sent a negative number")
				return
			}
		} else {
			log.Fatalf("Not a grpc error %v\n", err)
		}
	}
	log.Printf("A result is: %f", rc.Num)
}

func main() {
	c, err := grpc.Dial("localhost:50069", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	defer c.Close()

	client := pb.NewGetSQRTClient(c)

	getSQRT(client)
}
