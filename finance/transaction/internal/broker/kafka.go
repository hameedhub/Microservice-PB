// Reference:
// https://github.com/confluentinc/confluent-kafka-go doc

package broker

import (
	"context"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	AccountDeposit = "account_deposit"
	CreateTransfer = "create_transfer"
	TransferStatus = "transfer_status"
)

type KafkaClient struct {
	Server string
	Group  string
	Topics []Topic
}
type Topic struct {
	Topic             string
	NumPartitions     int
	ReplicationFactor int
	Priority          string
}

func NewKafkaClient(server, group string, topics []Topic) (*KafkaClient, error) {
	return &KafkaClient{
		Server: server,
		Group:  group,
		Topics: topics,
	}, nil

}

// connect admin client and create topics
func NewKafkaAdminClientCreateTopic(client *KafkaClient) error {
	a, err := kafka.NewAdminClient(&kafka.ConfigMap{"bootstrap.servers": client.Server})
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// make list of topics
	topics := make([]kafka.TopicSpecification, len(client.Topics))
	for i, topic := range client.Topics {
		t := kafka.TopicSpecification{
			Topic:             topic.Topic,
			NumPartitions:     int(topic.NumPartitions),
			ReplicationFactor: int(topic.ReplicationFactor),
		}
		//TODO: Implement check if topic already exist
		topics[i] = t
	}

	tps, err := a.CreateTopics(ctx, topics, kafka.SetAdminOperationTimeout(100))
	if err != nil {
		return err
	} else {
		for _, tp := range tps {
			fmt.Printf("%s\n", tp)
		}
	}

	return nil
}
