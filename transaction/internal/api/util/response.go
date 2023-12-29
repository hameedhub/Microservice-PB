package api

type Response struct {
	Status  int    `json:"status"`
	Data    string `json:"data"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type Payload struct {
	Amount  float64 `json:"amount"`
	Account int64   `json:"account"`
	Ref     int64   `json:"ref"`
	Status  string  `json:"status"`
}
