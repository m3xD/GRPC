// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: greet.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GreetServerClient is the client API for GreetServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetServerClient interface {
	Greet(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type greetServerClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServerClient(cc grpc.ClientConnInterface) GreetServerClient {
	return &greetServerClient{cc}
}

func (c *greetServerClient) Greet(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/greet.GreetServer/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetServerServer is the server API for GreetServer service.
// All implementations must embed UnimplementedGreetServerServer
// for forward compatibility
type GreetServerServer interface {
	Greet(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedGreetServerServer()
}

// UnimplementedGreetServerServer must be embedded to have forward compatible implementations.
type UnimplementedGreetServerServer struct {
}

func (UnimplementedGreetServerServer) Greet(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}
func (UnimplementedGreetServerServer) mustEmbedUnimplementedGreetServerServer() {}

// UnsafeGreetServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServerServer will
// result in compilation errors.
type UnsafeGreetServerServer interface {
	mustEmbedUnimplementedGreetServerServer()
}

func RegisterGreetServerServer(s grpc.ServiceRegistrar, srv GreetServerServer) {
	s.RegisterService(&GreetServer_ServiceDesc, srv)
}

func _GreetServer_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServerServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetServer/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServerServer).Greet(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GreetServer_ServiceDesc is the grpc.ServiceDesc for GreetServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetServer",
	HandlerType: (*GreetServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetServer_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greet.proto",
}