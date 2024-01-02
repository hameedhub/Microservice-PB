package api

import (
	"encoding/json"
	"microservice-pb/internal/broker"
	"microservice-pb/internal/domain"
	"net/http"
)

type transactionApi struct {
	repo        domain.TransactionRepository
	kafkaClient *broker.KafkaClient
}

type TransactionApi interface {
	Get(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
}

func NewTransaction(repo domain.TransactionRepository, kafkaClient *broker.KafkaClient) *transactionApi {
	return &transactionApi{repo: repo, kafkaClient: kafkaClient}
}

func (trans transactionApi) Get(w http.ResponseWriter, r *http.Request) {
	t := &[]domain.Transaction{}
	trans.repo.GetTransaction(1, t)
	transactions, _ := json.Marshal(t)

	w.Write([]byte(transactions))
}
func (trans transactionApi) Update(w http.ResponseWriter, r *http.Request) {
	t := &domain.Transaction{}
	trans.repo.UpdateTransaction(1, domain.Success, t)
	transaction, _ := json.Marshal(t)
	broker.Publish(trans.kafkaClient, "update_transaction", string(transaction))
}
