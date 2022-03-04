package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/juanmanuel-tirado/savetheworldwithgo/15_grpc/grpc/example_01/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (u *UserServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.User, error) {
	fmt.Println("Server received:", req.String())
	return &pb.User{UserId: "John", Email: "john@gmail.com"}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &UserServer{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
