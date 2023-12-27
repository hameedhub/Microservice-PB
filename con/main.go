package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func consumeMessages(consumer *kafka.Consumer) {
	run := true
	for run {
		ev := consumer.Poll(100)
		if ev == nil {
			continue
		}

		switch e := ev.(type) {
		case *kafka.Message:

			fmt.Printf("Received message on topic %s [%d] at offset %d: %s\n",
				*e.TopicPartition.Topic, e.TopicPartition.Partition, e.TopicPartition.Offset, string(e.Value))

		case kafka.Error:
			fmt.Printf("Error: %v\n", e)
			run = false // Terminate on error (change as per requirement)
		default:
			fmt.Printf("Ignored event: %v\n", e)
		}
	}
}

func main() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		fmt.Printf("Error creating consumer: %v\n", err)
		return
	}
	defer c.Close()

	topics := []string{"high", "medium", "low"}

	err = c.SubscribeTopics(topics, nil)
	if err != nil {
		fmt.Printf("Error subscribing to topics: %v\n", err)
		return
	}

	go func() {
		<-sigchan // Wait for signal to gracefully shutdown (e.g., CTRL+C)
		fmt.Println("Received termination signal. Closing consumer...")
		c.Close() // Close consumer on signal
	}()

	consumeMessages(c)
}
