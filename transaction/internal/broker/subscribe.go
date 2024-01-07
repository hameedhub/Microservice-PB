package broker

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"transaction/internal/domain"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Subscribe(client *KafkaClient, repo domain.TransactionRepository, topics []string) error {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": client.Server,
		"group.id":          client.Group,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return err
	}
	defer c.Close()

	err = c.SubscribeTopics(topics, nil)
	if err != nil {
		return err
	}

	go func() {
		<-sigchan
		fmt.Println("Termination signal. Closing consumer")
		c.Close()
	}()
	consume(c, repo)
	return nil
}

func consume(consumer *kafka.Consumer, repo domain.TransactionRepository) {
	run := true
	for run {
		ev := consumer.Poll(100)
		if ev == nil {
			continue
		}

		switch e := ev.(type) {
		case *kafka.Message:
			if *e.TopicPartition.Topic == AccountDeposit {
				ac := domain.Transaction{}
				r := strings.NewReader(string(e.Value))
				_ = json.NewDecoder(r).Decode(&ac)
				ac.Status = domain.Success
				repo.CreateTransaction(ac)
			}
		case kafka.Error:
			fmt.Printf("Error: %v\n", e)
			run = false // Terminate on error (change as per requirement)
		default:
			// fmt.Printf("Ignored event: %v\n", e)
		}
	}
}
