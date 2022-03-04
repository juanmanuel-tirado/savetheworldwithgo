package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/juanmanuel-tirado/savetheworldwithgo/15_grpc/interceptors/example_01/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (u *UserServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.User, error) {
	fmt.Println("Server received:", req.String())
	return &pb.User{UserId: "John", Email: "john@gmail.com"}, nil
}

func AuthServerInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	md, found := metadata.FromIncomingContext(ctx)
	if !found {
		return nil, status.Errorf(codes.InvalidArgument, "metadata not found")
	}
	password, found := md["password"]
	if !found {
		return nil, status.Errorf(codes.Unauthenticated, "password not found")
	}

	if password[0] != "go" {
		return nil, status.Errorf(codes.Unauthenticated, "password not valid")
	}

	h, err := handler(ctx, req)
	return h, err
}

func withAuthServerInterceptor() grpc.ServerOption {
	return grpc.UnaryInterceptor(AuthServerInterceptor)
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer(withAuthServerInterceptor())
	pb.RegisterUserServiceServer(s, &UserServer{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
