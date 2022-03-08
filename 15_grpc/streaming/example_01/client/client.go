package main

import (
	"context"
	"fmt"
	"io"
	"time"

	pb "github.com/juanmanuel-tirado/savetheworldwithgo/15_grpc/streaming/example_01/numbers"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewNumServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	stream, err := c.Rnd(ctx, &pb.NumRequest{N: 5, From: 0, To: 100})
	if err != nil {
		panic(err)
	}

	done := make(chan bool)
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				panic(err)
			}
			fmt.Println("Received:", resp.String())
		}
	}()
	<-done
	fmt.Println("Client done")
}
