package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"transaction/internal/broker"
	"transaction/internal/domain"
)

type transactionApi struct {
	repo        domain.TransactionRepository
	kafkaClient *broker.KafkaClient
}

type TransactionApi interface {
	Get(w http.ResponseWriter, r *http.Request)
}

func NewTransaction(repo domain.TransactionRepository, kafkaClient *broker.KafkaClient) *transactionApi {
	return &transactionApi{repo: repo, kafkaClient: kafkaClient}
}

func (trans transactionApi) Get(w http.ResponseWriter, r *http.Request) {
	account, _ := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])
	t := &[]domain.Transaction{}
	trans.repo.GetTransaction(int64(account), t)
	transactions, _ := json.Marshal(t)

	w.Write([]byte(transactions))
}
