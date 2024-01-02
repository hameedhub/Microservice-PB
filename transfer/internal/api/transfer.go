package api

import (
	"encoding/json"
	"net/http"
	"transfer/internal/broker"
	"transfer/internal/domain"
)

type transferApi struct {
	repo        domain.TransferRepository
	kafkaClient *broker.KafkaClient
}

type TransferApi interface {
	Create(w http.ResponseWriter, r *http.Request)
}

func NewTransfer(repo domain.TransferRepository, kafkaClient *broker.KafkaClient) *transferApi {
	return &transferApi{repo: repo, kafkaClient: kafkaClient}
}

func (trans transferApi) Create(w http.ResponseWriter, r *http.Request) {
	t := &domain.Transfer{}
	trans.repo.Create(int64(t.Amount))
	transactions, _ := json.Marshal(t)

	w.Write([]byte(transactions))
}
