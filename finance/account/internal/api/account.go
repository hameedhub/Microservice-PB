package api

import (
	"account/internal/broker"
	"account/internal/domain"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/exp/rand"
)

type accountApi struct {
	repo        domain.AccountRepository
	kafkaClient *broker.KafkaClient
}

type AccountApi interface {
	Create(w http.ResponseWriter, r *http.Request)
	Counter(w http.ResponseWriter, r *http.Request)
}

func NewTransfer(repo domain.AccountRepository, kafkaClient *broker.KafkaClient) *accountApi {
	return &accountApi{repo: repo, kafkaClient: kafkaClient}
}

func (account accountApi) Create(w http.ResponseWriter, r *http.Request) {

	t := domain.Account{}
	_ = json.NewDecoder(r.Body).Decode(&t)
	t.Status = domain.Active
	t.Account = int64(rand.Int63n(99999999))
	account.repo.Create(t)
	trans, _ := json.Marshal(t)
	broker.Publish(account.kafkaClient, broker.CreateAccount, string(trans))

	w.Write([]byte(trans))
}

func (account accountApi) Counter(w http.ResponseWriter, r *http.Request) {
	d := domain.Deposit{}
	_ = json.NewDecoder(r.Body).Decode(&d)
	account.repo.Update(d.Account, d.Amount, d.Type)
	dp, _ := json.Marshal(d)
	broker.Publish(account.kafkaClient, broker.AccountDeposit, string(dp))

	w.Write([]byte(dp))
}

func (account accountApi) Get(w http.ResponseWriter, r *http.Request) {
	accountNum, _ := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])
	ac := &domain.Account{}
	account.repo.Get(int64(accountNum), ac)
	accJson, _ := json.Marshal(ac)
	w.Write([]byte(accJson))

}
