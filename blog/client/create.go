package main

import (
	"context"
	pb "grpc/blog/proto"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {

	blogRequest := &pb.Blog{
		Author:  "Khanh",
		Title:   "First blog",
		Content: "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blogRequest)

	if err != nil {
		log.Fatal("Error")
	}

	log.Print("Blog has been created\n", res.Id)
	return res.Id
}
