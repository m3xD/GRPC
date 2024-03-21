package main

import (
	pb "grpc/blog/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	c, err := grpc.Dial("localhost:50069", grpc.WithInsecure())

	if err != nil {
		log.Fatalln("Error", err)
	}

	defer c.Close()

	client := pb.NewBlogServiceClient(c)

	createBlog(client)
}
