// Reference:
// https://github.com/confluentinc/confluent-kafka-go doc

package transport

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaClient struct {
	Server string
	Group  string
}

func NewKafkaClient(server, group string) (*KafkaClient, error) {
	return &KafkaClient{
		Server: server,
		Group:  group,
	}, nil

}

func NewProducer(client *KafkaClient) (*kafka.Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": client.Server,
	})
	if err != nil {
		return nil, err
	}
	defer p.Close()

	return p, nil
}

func NewConsumer(client *KafkaClient) (*kafka.Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": client.Server,
		"group.id":          client.Group,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, err
	}
	defer c.Close()

	return c, nil
}
