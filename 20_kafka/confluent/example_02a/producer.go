package main

import (
    "encoding/json"
    "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

type User struct {
    Name string `json:"name"`
    Email string `json:"email"`
}

func main() {
    cfg := kafka.ConfigMap{"bootstrap.servers": "localhost"}
    p, err := kafka.NewProducer(&cfg)
    if err != nil {
        panic(err)
    }
    defer p.Close()

    topic := "userTopic"
    users := []User{{"John", "john@gmail.com"},{"Mary", "mary@email.com"}}
    for _, u := range users {
        encoded, err := json.Marshal(u)
        if err != nil {
            panic(err)
        }
        p.Produce(&kafka.Message{
            TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
            Value: encoded,
        },nil)
    }
    p.Flush(1000)
}
