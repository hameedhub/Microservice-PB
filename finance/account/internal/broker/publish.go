package broker

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

// publish using produce
func Publish(client *KafkaClient, topic string, message string) {
	p, _ := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": client.Server})
	defer p.Close()
	var priority string
	for _, t := range client.Topics {
		if t.Topic == topic {
			priority = t.Priority
		}
	}
	produce(p, topic, message, priority)

	p.Flush(100)
}

// produce message ..
func produce(p *kafka.Producer, topic string, message, priority string) {
	fmt.Print(&topic)
	deliveryChan := make(chan kafka.Event)
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(priority),
		Value:          []byte(message),
		Timestamp:      time.Now(),
	}, deliveryChan)
}
