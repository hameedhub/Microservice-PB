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
			Topic:             "high",
			NumPartitions:     10,
			ReplicationFactor: 1},
			{Topic: "medium",
				NumPartitions:     4,
				ReplicationFactor: 1},
			{Topic: "low",
				NumPartitions:     2,
				ReplicationFactor: 1},
			{Topic: "default",
				//https://kafka.apache.org/081/documentation.html#topic-config
				// Requires even number else it defaulted to zero
				NumPartitions:     16,
				ReplicationFactor: 1},
		},
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

	priorityLevels := []string{"default", "high", "medium", "low"}
	for i := 0; i < 120; i++ {
		priority := priorityLevels[i%len(priorityLevels)]
		fmt.Printf("%s - %d\n", priority, i)
		produce(p, priority, payload)
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
