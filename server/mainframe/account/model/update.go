package model

type UpdateAccountInput struct {
	IBAN  string  `json:"iban"`
	Owner string  `json:"owner"`
	Cash  float32 `json:"cash"`
}
