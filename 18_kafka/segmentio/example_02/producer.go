package main

import (
    "context"
    "fmt"
    "github.com/segmentio/kafka-go"
)

func main() {
    w := &kafka.Writer{
        Addr:         kafka.TCP("localhost:9092"),
        Topic:        "helloTopic",
        Balancer:     &kafka.LeastBytes{},
    }
    defer w.Close()

    msgs := []string{"Save", "the", "world", "with", "Go!!!"}
    for _, m := range msgs {
        err := w.WriteMessages(context.Background(), kafka.Message{Value: []byte(m)})
        if err != nil {
            panic(err)
        }
        fmt.Printf("Sent message: %s\n", m)
    }
    fmt.Printf("Producer sent: %d bytes\n", w.Stats().Bytes)
}
