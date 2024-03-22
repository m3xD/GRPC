package main

import (
	"context"

	pb "grpc/blog/proto"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {

	data := BlogItem{
		AuthorId: in.Author,
		Title:    in.Author,
		Content:  in.Author,
	}

	res, err := collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Internal error",
		)
	}
	obid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"Error",
		)
	}
	return &pb.BlogId{
		Id: obid.Hex(),
	}, nil
}
