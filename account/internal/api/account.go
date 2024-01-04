package api

import (
	"account/internal/broker"
	"account/internal/domain"
	"encoding/json"
	"net/http"
)

type accountApi struct {
	repo        domain.AccountRepository
	kafkaClient *broker.KafkaClient
}

type AccountApi interface {
	Create(w http.ResponseWriter, r *http.Request)
}

func NewTransfer(repo domain.AccountRepository, kafkaClient *broker.KafkaClient) *accountApi {
	return &accountApi{repo: repo, kafkaClient: kafkaClient}
}

func (account accountApi) Create(w http.ResponseWriter, r *http.Request) {
	t := &domain.Account{}
	account.repo.Create(int64(t.Account))
	transactions, _ := json.Marshal(t)

	w.Write([]byte(transactions))
}

func (account accountApi) Deposit(w http.ResponseWriter, r *http.Request) {

}
