package main

import (
    "fmt"
    "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
    cfg := kafka.ConfigMap{"bootstrap.servers": "localhost:9092"}
    p, err := kafka.NewProducer(&cfg)
    if err != nil {
        panic(err)
    }
    defer p.Close()

    deliveryChan := make(chan kafka.Event,10)
    defer close(deliveryChan)
    topic := "helloTopic"
    msgs := []string{"Save", "the", "world", "with", "Go!!!"}
    for _, word := range msgs {
        err = p.Produce(&kafka.Message{
            TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
            Value: []byte(word),
        }, deliveryChan)
        if err != nil {
            panic(err)
        }
        e := <-deliveryChan
        m := e.(*kafka.Message)
        fmt.Printf("Sent %v\n",m)
    }
    // p.Flush(1000)
}
