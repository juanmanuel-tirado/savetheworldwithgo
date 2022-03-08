package main

import (
	"context"
	"fmt"
	"runtime"
	"time"

	pb "github.com/juanmanuel-tirado/savetheworldwithgo/15_grpc/interceptors/example_02/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func ClientLoggerInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption) error {

	os := runtime.GOOS
	zone, _ := time.Now().Zone()

	ctx = metadata.AppendToOutgoingContext(ctx, "os", os)
	ctx = metadata.AppendToOutgoingContext(ctx, "zone", zone)

	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}

func withUnaryClientLoggerInterceptor() grpc.DialOption {
	return grpc.WithUnaryInterceptor(ClientLoggerInterceptor)
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(),
		grpc.WithBlock(), withUnaryClientLoggerInterceptor())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetUser(ctx, &pb.UserRequest{UserId: "John"})
	if err != nil {
		panic(err)
	}
	fmt.Println("Client received:", r.String())
}
