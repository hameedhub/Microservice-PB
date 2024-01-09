package broker

import (
	"account/internal/domain"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Subscribe(client *KafkaClient, repo domain.AccountRepository, topics []string) error {
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
	consume(c, client, repo)
	return nil
}

func consume(consumer *kafka.Consumer, client *KafkaClient, repo domain.AccountRepository) {
	run := true
	for run {
		ev := consumer.Poll(100)
		if ev == nil {
			continue
		}

		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Println(*e.TopicPartition.Topic)
			if *e.TopicPartition.Topic == CreateTransfer {
				tf := domain.Transfer{}
				r := strings.NewReader(string(e.Value))
				_ = json.NewDecoder(r).Decode(&tf)
				// if account exist
				ac := &domain.Account{}
				ad := &domain.Account{}
				repo.Get(tf.CreditAccount, ac)
				repo.Get(tf.DebitAccount, ad)
				if ac.Account == 0 || ad.Account == 0 {
					tf.Status = domain.Failed
					f, _ := json.Marshal(tf)
					Publish(client, TransferStatus, string(f))
				}

			}

		case kafka.Error:
			fmt.Printf("Error: %v\n", e)
			run = false // Terminate on error (change as per requirement)
		default:
			// fmt.Printf("Ignored event: %v\n", e)
		}
	}
}
