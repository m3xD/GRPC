package main

import (
	"context"
	"fmt"
	pb "grpc/Max/proto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

func getMax(client pb.GetMaxClient) {

	rq := []int{1, 5, 3, 6, 2, 20}

	str, err := client.Max(context.Background())

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, num := range rq {
			str.Send(&pb.GetNumber{
				Num: int32(num),
			})
			time.Sleep(1 * time.Second)
		}
		str.CloseSend()
	}()

	go func() {
		for {
			msg, err := str.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error %v", err)
			}

			fmt.Print(msg.Num, " ")
		}
		close(waitc)
	}()
	<-waitc
}

func main() {
	c, err := grpc.Dial("localhost:50069", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	defer c.Close()

	client := pb.NewGetMaxClient(c)

	getMax(client)
}
