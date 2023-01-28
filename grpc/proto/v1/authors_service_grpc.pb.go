// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.3
// source: proto/v1/authors_service.proto

package v1

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

// AuthorsServiceClient is the client API for AuthorsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorsServiceClient interface {
	GetAuthor(ctx context.Context, in *GetAuthorRequest, opts ...grpc.CallOption) (*Author, error)
	ListAuthors(ctx context.Context, in *ListAuthorsRequest, opts ...grpc.CallOption) (*ListAuthorsResponse, error)
}

type authorsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorsServiceClient(cc grpc.ClientConnInterface) AuthorsServiceClient {
	return &authorsServiceClient{cc}
}

func (c *authorsServiceClient) GetAuthor(ctx context.Context, in *GetAuthorRequest, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, "/authors.v1.AuthorsService/GetAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorsServiceClient) ListAuthors(ctx context.Context, in *ListAuthorsRequest, opts ...grpc.CallOption) (*ListAuthorsResponse, error) {
	out := new(ListAuthorsResponse)
	err := c.cc.Invoke(ctx, "/authors.v1.AuthorsService/ListAuthors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorsServiceServer is the server API for AuthorsService service.
// All implementations must embed UnimplementedAuthorsServiceServer
// for forward compatibility
type AuthorsServiceServer interface {
	GetAuthor(context.Context, *GetAuthorRequest) (*Author, error)
	ListAuthors(context.Context, *ListAuthorsRequest) (*ListAuthorsResponse, error)
	mustEmbedUnimplementedAuthorsServiceServer()
}

// UnimplementedAuthorsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorsServiceServer struct {
}

func (UnimplementedAuthorsServiceServer) GetAuthor(context.Context, *GetAuthorRequest) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthor not implemented")
}
func (UnimplementedAuthorsServiceServer) ListAuthors(context.Context, *ListAuthorsRequest) (*ListAuthorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAuthors not implemented")
}
func (UnimplementedAuthorsServiceServer) mustEmbedUnimplementedAuthorsServiceServer() {}

// UnsafeAuthorsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorsServiceServer will
// result in compilation errors.
type UnsafeAuthorsServiceServer interface {
	mustEmbedUnimplementedAuthorsServiceServer()
}

func RegisterAuthorsServiceServer(s grpc.ServiceRegistrar, srv AuthorsServiceServer) {
	s.RegisterService(&AuthorsService_ServiceDesc, srv)
}

func _AuthorsService_GetAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorsServiceServer).GetAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authors.v1.AuthorsService/GetAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorsServiceServer).GetAuthor(ctx, req.(*GetAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorsService_ListAuthors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAuthorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorsServiceServer).ListAuthors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authors.v1.AuthorsService/ListAuthors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorsServiceServer).ListAuthors(ctx, req.(*ListAuthorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthorsService_ServiceDesc is the grpc.ServiceDesc for AuthorsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthorsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authors.v1.AuthorsService",
	HandlerType: (*AuthorsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuthor",
			Handler:    _AuthorsService_GetAuthor_Handler,
		},
		{
			MethodName: "ListAuthors",
			Handler:    _AuthorsService_ListAuthors_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/v1/authors_service.proto",
}
