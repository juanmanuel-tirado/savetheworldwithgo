package main

import (
    "fmt"
    "github.com/confluentinc/confluent-kafka-go/kafka"
)

const COMMIT_N = 2

func Committer(c *kafka.Consumer) {
    offsets, err := c.Commit()
    if err != nil {
        fmt.Printf("Error: %s\n", err)
        return
    }
    fmt.Printf("Offset: %#v\n",offsets[0].String())
}

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
    counter := 0
    for {
        ev := c.Poll(1000)
        switch e := ev.(type) {
        case *kafka.Message:
            counter += 1
            if counter % COMMIT_N == 0 {
                go Committer(c)
            }
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
