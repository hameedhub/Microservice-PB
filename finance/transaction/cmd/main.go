package main

import (
	"log"
	"net/http"
	"time"
	"transaction/internal/api"
	"transaction/internal/broker"
	"transaction/internal/config"
	"transaction/internal/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Read configuration data ...
	c, err := config.ReadConfig(".")
	if err != nil {
		panic(err)
	}

	// new logger
	l, err := config.NewLogger("logs")
	//setup database ORM sqlite DB
	// REF: https://gorm.io/docs/
	db, err := gorm.Open(sqlite.Open("transactions.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// // Schema migration
	db.AutoMigrate(&domain.Transaction{})

	// topics  REF: https://kafka.apache.org/documentation/#topicconfigs
	topics := []broker.Topic{
		{Topic: "update_transaction", NumPartitions: int(c.MEDIUM_PRIORITY_PARTITION), ReplicationFactor: 1, Priority: "MEDIUM_PRIORITY_PARTITION"},
	}

	// create client
	client, err := broker.NewKafkaClient(c.KAFKA_SERVER, c.KAFKA_CONSUMER_GROUP, topics)
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
	controller := api.NewTransaction(repo, client)

	// http REF: https://pkg.go.dev/net/http
	mux := http.NewServeMux()
	mux.HandleFunc("/transaction/", func(w http.ResponseWriter, r *http.Request) {
		// get transactions
		if r.Method == http.MethodGet {
			controller.Get(w, r)
		}
	})

	// server configuration
	server := http.Server{
		IdleTimeout:  time.Duration(c.SERVER_IDLE_TIMEOUT) * time.Second,
		ReadTimeout:  time.Duration(c.SERVER_READ_TIMEOUT) * time.Second,
		WriteTimeout: time.Duration(c.SERVER_WRITE_TIMEOUT) * time.Second,
		Addr:         c.SERVER_PORT,
		Handler:      mux,
	}

	// listen to topics
	go broker.Subscribe(client, repo, []string{broker.AccountDeposit, broker.CreateTransfer, broker.TransferStatus, broker.CreateAccount}, l)

	// listen to http requests
	log.Fatal(server.ListenAndServe())

}
