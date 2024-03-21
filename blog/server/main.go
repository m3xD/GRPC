package main

import (
	"context"
	"fmt"
	pb "grpc/blog/proto"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type server struct {
	pb.BlogServiceServer
}

var collection *mongo.Collection

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))

	if err != nil {
		log.Fatalf("An error occur %v", err)
	}

	err = client.Connect(context.Background())

	if err != nil {
		log.Fatalf("An error occur %v", err)
	}

	collection = client.Database("blogdb").Collection("blog")

	addr := "localhost:50069"
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("An error occur %v", err)
	}

	fmt.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()

	pb.RegisterBlogServiceServer(s, &server{})

	err = s.Serve(lis)

	if err != nil {
		log.Fatalf("An error occur %v", err)
	}

}
