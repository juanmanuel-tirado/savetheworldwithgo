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

    conn.SetWriteDeadline(time.Now().Add(3*time.Second))
    msgs := []string{"Save", "the", "world", "with", "Go!!!"}
    for _, m := range msgs {
        l, err := conn.WriteMessages(kafka.Message{Value: []byte(m)})
        if err != nil {
            panic(err)
        }
        fmt.Printf("Sent %d bytes: %s\n", l,m)
    }
}
