package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net"
	"time"

	pb "github.com/juanmanuel-tirado/savetheworldwithgo/15_grpc/streaming/example_01/numbers"
	"google.golang.org/grpc"
)

type NumServer struct {
	pb.UnimplementedNumServiceServer
}

func (n *NumServer) Rnd(req *pb.NumRequest, stream pb.NumService_RndServer) error {
	fmt.Println(req.String())
	if req.N <= 0 {
		return errors.New("N must be greater than zero")
	}
	if req.To <= req.From {
		return errors.New("to must be greater or equal than from")
	}
	done := make(chan bool)
	go func() {
		for counter := 0; counter < int(req.N); counter++ {
			i := rand.Intn(int(req.To)-int(req.From)+1) + int(req.To)
			resp := pb.NumResponse{I: int64(i), Remaining: req.N - int64(counter)}
			stream.Send(&resp)
			time.Sleep(time.Second)
		}
		done <- true
	}()
	<-done
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()

	pb.RegisterNumServiceServer(s, &NumServer{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
