package main

import (
    "context"
    "fmt"
    pb "github.com/juanmanuel-tirado/savetheworldwithgo/13_grpc/interceptors/example_02/user"
    "google.golang.org/grpc"
    "google.golang.org/grpc/metadata"
    "net"
)

type UserServer struct{
    pb.UnimplementedUserServiceServer
}

func (u *UserServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.User, error) {
    fmt.Println("Server received:", req.String())
    return &pb.User{UserId: "John", Email: "john@gmail.com"}, nil
}

func RequestLoggerInterceptor(ctx context.Context,
    req interface{},
    info *grpc.UnaryServerInfo,
    handler grpc.UnaryHandler) (interface{}, error){
        md, found := metadata.FromIncomingContext(ctx)
        if found {
            os, _ := md["os"]
            zone, _ := md["zone"]
            fmt.Printf("Request from %s using %s\n", zone, os)
        }

        h, err := handler(ctx, req)
        return h, err
}

func withRequestLoggerInterceptor() grpc.ServerOption {
    return grpc.UnaryInterceptor(RequestLoggerInterceptor)
}

func main() {
    lis, err := net.Listen("tcp", "localhost:50051")
    if err != nil {
        panic(err)
    }
    s := grpc.NewServer(withRequestLoggerInterceptor())
    pb.RegisterUserServiceServer(s, &UserServer{})


    if err := s.Serve(lis); err != nil {
        panic(err)
    }
}
