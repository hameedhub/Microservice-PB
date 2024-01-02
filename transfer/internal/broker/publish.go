package broker

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// publish using produce
func Publish(client *KafkaClient, topic string, message string) {
	p, _ := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": client.Server})

	defer p.Close()
	produce(p, topic, message)

	p.Flush(100)

}

// produce messsage
func produce(p *kafka.Producer, topic string, message string) {
	deliveryChan := make(chan kafka.Event)
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, deliveryChan)
}
