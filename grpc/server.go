package main

import (
	"context"
	v1 "github.com/songtomtom/graphql-mesh-gateway/grpc/proto/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	v1.UnimplementedAuthorsServiceServer
}

func (s *server) GetAuthor(ctx context.Context, in *v1.GetAuthorRequest) (*v1.Author, error) {
	return &v1.Author{
		Id:     in.GetId(),
		Name:   "dummy_name",
		Editor: "dummy_editor",
	}, nil
}

func (s *server) ListAuthors(ctx context.Context, in *v1.ListAuthorsRequest) (*v1.ListAuthorsResponse, error) {
	return &v1.ListAuthorsResponse{Items: []*v1.Author{}}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":3003")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	v1.RegisterAuthorsServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
