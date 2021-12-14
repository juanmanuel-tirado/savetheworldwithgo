package main

import (
    "fmt"
    "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
    c, err := kafka.NewConsumer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "group.id":          "helloGroup",
        "auto.offset.reset": "earliest",
    })
    if err != nil {
        panic(err)
    }
    defer c.Close()

    c.Subscribe("helloTopic", nil)
    for {
        ev := c.Poll(1000)
        switch e := ev.(type) {
        case *kafka.Message:
            c.Commit()
            fmt.Printf("Msg on %s: %s\n", e.TopicPartition, string(e.Value))
        case kafka.PartitionEOF:
            fmt.Printf("%v\n",e)
        case kafka.Error:
            fmt.Printf("Error: %v\n", e)
            break
        default:
            fmt.Printf("Ignored: %v\n",e)
        }
    }
}
