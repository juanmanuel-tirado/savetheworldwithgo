package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/juanmanuel-tirado/savetheworldwithgo/15_grpc/grpc/example_01/user"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
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
