package model

type Price struct {
	Amount   float32 `json:"amount"`
	Currency string  `json:"currency"`
}
