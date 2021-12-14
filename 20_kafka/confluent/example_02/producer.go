package main

import (
    "fmt"
    "github.com/confluentinc/confluent-kafka-go/kafka"
)

func Handler(c chan kafka.Event) {
    for {
        e := <- c
        if e == nil {
            return
        }
        m := e.(*kafka.Message)
        if m.TopicPartition.Error != nil {
            fmt.Printf("Partition error %s\n",m.TopicPartition.Error)
        } else {
            fmt.Printf("Sent %v: %s\n",m,string(m.Value))
        }
    }
}

func main() {
    cfg := kafka.ConfigMap{"bootstrap.servers": "localhost:9092"}
    p, err := kafka.NewProducer(&cfg)
    if err != nil {
        panic(err)
    }
    defer p.Close()

    delivery_chan := make(chan kafka.Event,10)
    defer close(delivery_chan)
    go Handler(delivery_chan)

    topic := "helloTopic"
    msgs := []string{"Save", "the", "world", "with", "Go!!!"}
    for _, word := range msgs {
        err = p.Produce(&kafka.Message{
            TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
            Value: []byte(word),
        },delivery_chan)
    }
    p.Flush(10000)
}
