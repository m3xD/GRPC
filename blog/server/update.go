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

func (*server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {

	tmp, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"Not found by ID")
	}

	data := &BlogItem{
		AuthorId: in.Author,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": tmp},
		bson.M{"$set": data})

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Cannot update")
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Cannot find by ID")
	}

	return &emptypb.Empty{}, nil
}
