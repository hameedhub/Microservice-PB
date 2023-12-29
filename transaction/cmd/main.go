package main

import (
	"context"
	"log"
	"microservice-pb/internal/config"
	"microservice-pb/internal/domain"
	"microservice-pb/internal/transport"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Read configuration data ...
	config, err := config.ReadConfig(".")
	if err != nil {
		panic(err)
	}

	//setup database ORM
	// REF: https://gorm.io/docs/
	db, err := gorm.Open(sqlite.Open("transactions.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Schema migration
	db.AutoMigrate(&domain.Transaction{})

	// topics  REF: https://kafka.apache.org/documentation/#topicconfigs
	topics := []transport.Topic{
		{Topic: "high", NumPartitions: int(config.HIGH_PRIORITY_PARTITION), ReplicationFactor: 1},
		{Topic: "medium", NumPartitions: int(config.MEDIUM_PRIORITY_PARTITION), ReplicationFactor: 1},
		{Topic: "low", NumPartitions: int(config.LOW_PRIORITY_PARTITION), ReplicationFactor: 1},
	}

	// create client
	client, err := transport.NewKafkaClient(config.KAFKA_SERVER, config.KAFKA_CONSUMER_GROUP, topics)
	if err != nil {
		panic(err)
	}

	// create topic
	err = transport.NewKafkaAdminClientCreateTopic(client)
	if err != nil {
		panic(err)
	}

	// http REF: https://pkg.go.dev/net/http
	mux := http.NewServeMux()
	mux.HandleFunc("/transaction", func(w http.ResponseWriter, r *http.Request) {
		// create transaction
		// if r.Method == http.MethodPost {

		// }
		// if r.Method == http.MethodGet {

		// }
	})

	// configure server
	server := http.Server{
		IdleTimeout:  time.Duration(config.SERVER_IDLE_TIMEOUT) * time.Second,
		ReadTimeout:  time.Duration(config.SERVER_READ_TIMEOUT) * time.Second,
		WriteTimeout: time.Duration(config.SERVER_WRITE_TIMEOUT) * time.Second,
		Addr:         config.SERVER_PORT,
		Handler:      mux,
	}

	// run server
	go func() {
		log.Printf("Server is listening to port %v", config.SERVER_PORT)
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// gracefull shutdown
	signChan := make(chan os.Signal)
	signChan <- os.Kill
	signChan <- os.Interrupt
	ctx, er := context.WithTimeout(context.Background(), 20*time.Second)
	if er != nil {
		log.Fatal(er)
	}
	server.Shutdown(ctx)

}
