package model

type Payment struct {
	Id      string         `json:"id" bson:"_id"`
	Type    string         `json:"type" bson:"type"`
	Value   PaymentValue   `json:"value" bson:"value"`
	Payer   Payer          `json:"payer" bson:"payer"`
	Payee   Payee          `json:"payee" bson:"payee"`
	Details string         `json:"details,omitempty" bson:"details,omitempty"`
	Outcome PaymentOutcome `json:"outcome,omitempty" bson:"outcome,omitempty"`
}

type PaymentValue struct {
	Amount   float32 `json:"amount" bson:"amount"`
	Currency string  `json:"currency" bson:"currency"`
}

type Payer struct {
	AccountIdentification AccountIdentification `json:"accountIdentification" bson:"accountIdentification"`
}

type Payee struct {
	Name                  string                `json:"name" bson:"name"`
	AccountIdentification AccountIdentification `json:"accountIdentification" bson:"accountIdentification"`
}

type AccountIdentification struct {
	Type  string `json:"type" bson:"type"`
	Value string `json:"value" bson:"value"`
}

type PaymentOutcome struct {
	Status    string `json:"status" bson:"status"`
	Message   string `json:"message" bson:"message"`
	Timestamp string `json:"ts" bson:"ts"`
}

type InsertPayment struct {
	Type    string       `json:"type" bson:"type"`
	Value   PaymentValue `json:"value" bson:"value"`
	Payer   Payer        `json:"payer" bson:"payer"`
	Payee   Payee        `json:"payee" bson:"payee"`
	Details string       `json:"details" bson:"details"`
}
