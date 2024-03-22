package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "grpc/blog/proto"
)

func (*server) ListBlog(e *emptypb.Empty, stream pb.BlogService_ListBlogServer) error {

	cur, err := collection.Find(context.Background(), primitive.D{{}})

	if err != nil {
		return status.Errorf(
			codes.Internal,
			"Error")
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err = cur.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal,
				"Error")
		}

		stream.Send(documentToBlog(data))
	}

	err = cur.Err()

	if err != nil {
		return status.Errorf(
			codes.Internal,
			"Error")
	}
	return nil
}
