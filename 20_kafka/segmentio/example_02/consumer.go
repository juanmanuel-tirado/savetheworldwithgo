package main

import (
    "context"
    "fmt"
    "github.com/segmentio/kafka-go"
)

func main () {
    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers:    []string{"localhost:9092"},
        Partition:  0,
        Topic:      "helloTopic",
        GroupID: "testGroup",
        MinBytes:   10e3,
        MaxBytes:   10e6,
    })

    defer r.Close()

    for {
        m, err := r.FetchMessage(context.Background())
        if err != nil {
            break
        }
        if err := r.CommitMessages(context.Background(), m); err != nil {
            panic(err)
        }
        fmt.Printf("Topic %s msg: %s\n", m.Topic, m.Value)
    }
}
