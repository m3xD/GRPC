package main

import (
	"context"
	pb "grpc/blog/proto"
	"log"
)

func deleteBlog(c pb.BlogServiceClient, id string) {

	rq := &pb.BlogId{
		Id: id,
	}

	_, err := c.DeleteBlog(context.Background(), rq)
	if err != nil {
		log.Fatalln("Error")
	}
	log.Println("Delete successfully")
}
