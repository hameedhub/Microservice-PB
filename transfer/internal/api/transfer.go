package api

import (
	"encoding/json"
	"math/rand"
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
	json.NewDecoder(r.Body).Decode(t)
	t.Status = domain.Pending
	t.Ref = int64(rand.Int63n(99999))
	trans.repo.Create(t)
	transactions, _ := json.Marshal(t)
	broker.Publish(trans.kafkaClient, broker.CreateTransfer, string(transactions))

	w.Write([]byte(transactions))
}
