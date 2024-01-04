package main

import (
	"account/internal/api"
	"account/internal/broker"
	"account/internal/config"
	"account/internal/domain"
	"log"
	"net/http"
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

	//setup database ORM sqlite DB
	// REF: https://gorm.io/docs/
	db, err := gorm.Open(sqlite.Open("accounts.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// // Schema migration
	db.AutoMigrate(&domain.Account{})

	// topics  REF: https://kafka.apache.org/documentation/#topicconfigs
	topics := []broker.Topic{
		{Topic: "create_account", NumPartitions: int(config.LOW_PRIORITY_PARTITION), ReplicationFactor: 1},
		{Topic: "account_deposit", NumPartitions: int(config.MEDIUM_PRIORITY_PARTITION), ReplicationFactor: 1},
	}

	// create client
	client, err := broker.NewKafkaClient(config.KAFKA_SERVER, config.KAFKA_CONSUMER_GROUP, topics)
	if err != nil {
		log.Fatal(err)
	}

	// // create topics
	err = broker.NewKafkaAdminClientCreateTopic(client)
	if err != nil {
		log.Fatal(err)
	}

	// // setting up repo for the controller to access db
	repo := domain.NewRepo(db)
	controller := api.NewTransfer(repo, client)

	// http REF: https://pkg.go.dev/net/http
	mux := http.NewServeMux()
	mux.HandleFunc("/account", func(w http.ResponseWriter, r *http.Request) {
		// update transfer
		if r.Method == http.MethodPost {
			controller.Create(w, r)
		}
		if r.Method == http.MethodPatch {
			controller.Deposit(w, r)
		}
	})

	// server configuration
	server := http.Server{
		IdleTimeout:  time.Duration(config.SERVER_IDLE_TIMEOUT) * time.Second,
		ReadTimeout:  time.Duration(config.SERVER_READ_TIMEOUT) * time.Second,
		WriteTimeout: time.Duration(config.SERVER_WRITE_TIMEOUT) * time.Second,
		Addr:         config.SERVER_PORT,
		Handler:      mux,
	}

	// listen to topics
	go broker.Subscribe(client, []string{"update_transaction"})

	// listen to http requests
	log.Fatal(server.ListenAndServe())

}
