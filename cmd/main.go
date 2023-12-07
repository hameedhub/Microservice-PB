// https://betterprogramming.pub/how-are-you-structuring-your-go-microservices-a355d6293932
package main

import (
	"microservice-pb/internal/config"
	"microservice-pb/internal/transport"
)

func main() {
	// read config data
	config, err := config.ReadConfig(".")

	if err != nil {
		panic(err)
	}

	_, err = transport.NewKafkaClient(config.KAFKA_SERVER, config.KAFKA_CONSUMER_GROUP)
	if err != nil {
		panic(err)
	}

}
