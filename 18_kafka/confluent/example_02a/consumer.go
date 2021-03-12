package main

import (
    "fmt"
    "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
    "encoding/json"
)

type User struct {
    Name string `json:"name"`
    Email string `json:"email"`
}

func main() {

    c, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost",
        "group.id":          "userGroup",
        "auto.offset.reset": "earliest",
    })
    if err != nil {
        panic(err)
    }
    defer c.Close()

    c.Subscribe("userTopic", nil)
    for {

        msg, err := c.ReadMessage(-1)
        if err != nil {
            fmt.Printf("Consumer error: #{err} (#{msg}\n)")
        }
        var u User
        err = json.Unmarshal(msg.Value, &u)
        if err == nil {
            fmt.Printf("Message on %s: %s\n", msg.TopicPartition, u)
        } else {
            // The client will automatically try to recover from all errors.
            fmt.Printf("Consumer error: %v (%v)\n", err, msg)
        }
    }
}
