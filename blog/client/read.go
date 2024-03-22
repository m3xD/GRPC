package main

import (
	"context"
	pb "grpc/blog/proto"
	"log"
)

func readBlog(c pb.BlogServiceClient, id string) {
	rq := &pb.BlogId{
		Id: id,
	}

	res, err := c.ReadBlog(context.Background(), rq)

	if err != nil {
		log.Fatalln("Error while reading")
	}

	log.Printf("Result is %v\n", res)
}
