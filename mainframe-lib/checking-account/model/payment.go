package model

type Payment struct {
	Id      string       `json:"id" bson:"_id"`
	Type    string       `json:"type" bson:"type"`
	Value   PaymentValue `json:"value" bson:"value"`
	Payer   Payer        `json:"payer" bson:"payer"`
	Payee   Payee        `json:"payee" bson:"payee"`
	Details string       `json:"details,omitempty" bson:"details,omitempty"`
}

type PaymentValue struct {
	Amount   float32 `json:"amount" bson:"amount"`
	Currency string  `json:"currency" bson:"currency"`
}

type Payer struct {
	Account string `json:"account" bson:"account"`
}

type Payee struct {
	Name                  string                `json:"name" bson:"name"`
	AccountIdentification AccountIdentification `json:"accountIdentification" bson:"accountIdentification"`
}

type AccountIdentification struct {
	Type  string `json:"type" bson:"type"`
	Value string `json:"value" bson:"value"`
}

type InsertPayment struct {
	Type    string       `json:"type" bson:"type"`
	Value   PaymentValue `json:"value" bson:"value"`
	Payer   Payer        `json:"payer" bson:"payer"`
	Payee   Payee        `json:"payee" bson:"payee"`
	Details string       `json:"details" bson:"details"`
}
