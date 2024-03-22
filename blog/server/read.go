package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "grpc/blog/proto"
)

func (*server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {

	res, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Cannot parse ID",
		)
	}
	data := &BlogItem{}
	filter := bson.M{"_id": res}

	ans := collection.FindOne(ctx, filter)

	err = ans.Decode(data)

	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"Not found ID")
	}
	return documentToBlog(data), nil
}
