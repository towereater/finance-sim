package model

type CheckingAccount struct {
	Id    string        `json:"id" bson:"_id"`
	IBAN  string        `json:"iban" bson:"iban"`
	Owner string        `json:"owner" bson:"owner"`
	Value CheckingValue `json:"value" bson:"value"`
}

type CheckingValue struct {
	Amount   float32 `json:"amount" bson:"amount"`
	Currency string  `json:"currency" bson:"currency"`
}

type InsertCheckingAccountInput struct {
	Owner string `json:"owner"`
}
