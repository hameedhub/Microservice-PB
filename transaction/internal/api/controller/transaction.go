package controller

import (
	"microservice-pb/internal/domain"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type transactionController struct {
	repo     domain.TransactionRepository
	producer *kafka.Producer
}

type TransactionController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

func NewTransactionController(repo domain.TransactionRepository, producer *kafka.Producer) *transactionController {
	return &transactionController{repo: repo, producer: producer}
}

func (trans transactionController) Create(w http.ResponseWriter, r *http.Request) {

}
func (trans transactionController) Get(w http.ResponseWriter, r *http.Request) {

}
