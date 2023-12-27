package main

import (
	"context"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// create client
	a, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		fmt.Printf("Failed to create Admin client: %s\n", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create topic
	results, err := a.CreateTopics(
		ctx, []kafka.TopicSpecification{{
			Topic:             "high",
			NumPartitions:     10,
			ReplicationFactor: 1},
			{Topic: "medium",
				NumPartitions:     5,
				ReplicationFactor: 1},
			{Topic: "low",
				NumPartitions:     2,
				ReplicationFactor: 1}},
		// Admin options
		kafka.SetAdminOperationTimeout(100))

	if err != nil {
		fmt.Printf("Failed to create topics: %v\n", err)
	} else {
		// Check results for more information
		for _, result := range results {
			fmt.Printf("%s\n", result)
		}
	}
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		fmt.Printf("Error creating producer: %v\n", err)
	}
	defer p.Close()

	produce(p, "low")
	produce(p, "medium")
	produce(p, "high")
	produce(p, "medium")
	produce(p, "high")
	produce(p, "high")
	p.Flush(100)
}

func produce(p *kafka.Producer, topic string) {

	deliveryChan := make(chan kafka.Event)

	err := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte("Hello " + topic),
	}, deliveryChan)

	if err != nil {
		fmt.Printf("Error producing message: %v\n", err)
	}

}
