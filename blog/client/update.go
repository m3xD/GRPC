package main

import (
	"context"
	pb "grpc/blog/proto"
	"log"
)

func updateBlog(c pb.BlogServiceClient, id string) {

	rq := &pb.Blog{
		Id:      id,
		Author:  "Not DK",
		Title:   "X",
		Content: "Y",
	}

	_, err := c.UpdateBlog(context.Background(), rq)

	if err != nil {
		log.Fatalln("Cannot update")
	}
}
