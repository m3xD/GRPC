package main

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "grpc/blog/proto"
	"io"
	"log"
)

func listBlog(c pb.BlogServiceClient) {

	res, err := c.ListBlog(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalln("Error")
	}

	for {
		msg, err := res.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln("Error")
		}

		log.Fatalf("Blog %v:\n", msg)
	}
}
