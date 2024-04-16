package broker

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
	"transaction/internal/config"
	"transaction/internal/domain"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Subscribe(client *KafkaClient, repo domain.TransactionRepository, topics []string, logger config.Logger) error {
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
	consume(c, repo, logger)
	return nil
}

func consume(consumer *kafka.Consumer, repo domain.TransactionRepository, logger config.Logger) {
	run := true
	for run {
		ev := consumer.Poll(1000)
		if ev == nil {
			continue
		}

		switch e := ev.(type) {
		case *kafka.Message:
			if *e.TopicPartition.Topic == AccountDeposit {
				ac := domain.Transaction{}
				r := strings.NewReader(string(e.Value))
				_ = json.NewDecoder(r).Decode(&ac)
				logger.Log(config.Log{
					Topic:        *e.TopicPartition.Topic,
					Priority:     string(e.Key),
					Service:      "account",
					SentTime:     e.Timestamp,
					ReceivedTime: time.Now(),
					PayloadSize:  len(string(e.Value)),
				})
				ac.Status = domain.Success
				repo.CreateTransaction(ac)
			}
			if *e.TopicPartition.Topic == CreateTransfer {
				tf := domain.Transfer{}
				r := strings.NewReader(string(e.Value))
				json.NewDecoder(r).Decode(&tf)

				logger.Log(config.Log{
					Topic:        *e.TopicPartition.Topic,
					Priority:     string(e.Key),
					Service:      "transfer",
					SentTime:     e.Timestamp,
					ReceivedTime: time.Now(),
					PayloadSize:  len(string(e.Value)),
				})
				// credit account
				ac := domain.Transaction{
					Amount:  tf.Amount,
					Account: tf.CreditAccount,
					Type:    domain.Credit,
					Ref:     tf.Ref,
					Status:  domain.Pending,
				}
				repo.CreateTransaction(ac)
				// debit account
				ad := domain.Transaction{
					Amount:  tf.Amount,
					Account: tf.DebitAccount,
					Type:    domain.Debit,
					Ref:     tf.Ref,
					Status:  domain.Pending,
				}
				repo.CreateTransaction(ad)
			}
			fmt.Println(*e.TopicPartition.Topic)
			if *e.TopicPartition.Topic == TransferStatus {
				tf := domain.Transfer{}
				r := strings.NewReader(string(e.Value))
				json.NewDecoder(r).Decode(&tf)
				logger.Log(config.Log{
					Topic:        *e.TopicPartition.Topic,
					Priority:     string(e.Key),
					Service:      "account",
					SentTime:     e.Timestamp,
					ReceivedTime: time.Now(),
					PayloadSize:  len(string(e.Value)),
				})

				repo.UpdateTransaction(tf.Ref, tf.Status)
			}
			if *e.TopicPartition.Topic == CreateAccount {
				logger.Log(config.Log{
					Topic:        *e.TopicPartition.Topic,
					Priority:     string(e.Key),
					Service:      "account",
					SentTime:     e.Timestamp,
					ReceivedTime: time.Now(),
					PayloadSize:  len(string(e.Value)),
				})
			}

		case kafka.Error:
			fmt.Printf("Error: %v\n", e)
			run = false // Terminate on error (change as per requirement)
		default:
			// fmt.Printf("Ignored event: %v\n", e)
		}
	}
}
