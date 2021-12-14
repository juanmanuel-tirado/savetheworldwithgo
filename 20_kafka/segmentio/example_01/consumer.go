package main

import (
    "context"
    "fmt"
    "github.com/segmentio/kafka-go"
    "time"
)

func main() {
    topic := "helloTopic"
    partition := 0
    conn, err := kafka.DialLeader(context.Background(),
        "tcp", "localhost:9092",topic,partition)
    if err != nil {
        panic(err)
    }

    defer conn.Close()

    conn.SetReadDeadline(time.Now().Add(time.Second))
    batch := conn.ReadBatch(10e3, 10e6)
    defer batch.Close()
    for {
        b := make([]byte, 10e3)
        l, err := batch.Read(b)
        if err != nil {
            break
        }
        fmt.Printf("Received %d: %s\n",l,string(b))
    }
}
