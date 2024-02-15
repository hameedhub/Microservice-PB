package main

import (
	"context"
	"fmt"
	"time"

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
			Topic:             "default",
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

	payload := "{\n    \"credit_account\": 38371524,\n    \"debit_account\": 36581830,\n    \"amount\":200\n}"

	for i := 0; i < 90; i++ {
		produce(p, "default", payload)
	}

	fmt.Println("Done............")

	p.Flush(100)
}

func produce(p *kafka.Producer, topic string, data string) {

	deliveryChan := make(chan kafka.Event)

	err := p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(topic),
		Value:          []byte(data),
		Timestamp:      time.Now(),
	}, deliveryChan)

	if err != nil {
		fmt.Printf("Error producing message: %v\n", err)
	}

}
