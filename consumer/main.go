package main

import (
	"consumer/csvlog"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func consumeMessages(consumer *kafka.Consumer, logger csvlog.Logger) {
	run := true
	for run {
		ev := consumer.Poll(500)
		if ev == nil {
			continue
		}

		switch e := ev.(type) {
		case *kafka.Message:
			logger.Log(csvlog.Log{
				Priority:     *e.TopicPartition.Topic,
				Partition:    e.TopicPartition.Partition,
				SentTime:     e.Timestamp,
				ReceivedTime: time.Now(),
				PayloadSize:  len(string(e.Value)),
			})
			//fmt.Printf("Received message on topic %s [%d] at offset %d: %s\n",
			//	*e.TopicPartition.Topic, e.TopicPartition.Partition, e.TopicPartition.Offset, string(e.Value))

		case kafka.Error:
			fmt.Printf("Error: %v\n", e)
			run = false // Terminate on error (change as per requirement)
		default:
			fmt.Printf("Ignored event: %v\n", e)
		}
	}
}

func main() {

	logger, err := csvlog.NewLogger("logs")
	if err != nil {
		fmt.Println("Error from logger")
	}

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

	topics := []string{"default"}
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

	consumeMessages(c, logger)
}
