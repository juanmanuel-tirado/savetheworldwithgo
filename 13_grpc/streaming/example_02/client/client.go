package main

import (
    "context"
    "fmt"
    pb "github.com/juanmanuel-tirado/savetheworldwithgo/13_grpc/streaming/example_02/numbers"
    "google.golang.org/grpc"
    "time"
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

    stream, err := c.Sum(ctx)
    if err != nil {
        panic(err)
    }

    from, to := 1,100

    for i:=from;i<=to;i++ {
        err = stream.Send(&pb.NumRequest{X:int64(i)})
        if err!= nil {
            panic(err)
        }
    }
    fmt.Println("Waiting for response...")
    result, err := stream.CloseAndRecv()
    if err != nil {
        panic(err)
    }
    fmt.Printf("The sum from %d to %d is %d\n", from,to, result.Total)
}

