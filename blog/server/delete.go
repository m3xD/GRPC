package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "grpc/blog/proto"
)

func (*server) DeleteBlog(ctx context.Context, blog *pb.BlogId) (*emptypb.Empty, error) {

	objectID, err := primitive.ObjectIDFromHex(blog.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Something happend")
	}

	filter := bson.M{"_id": objectID}

	c, err := collection.DeleteOne(ctx, &filter)

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Something happend")
	}

	if c.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Something happend")
	}

	return &emptypb.Empty{}, nil
}
