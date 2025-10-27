package model

type CheckingAccount struct {
	Id           string            `json:"id" bson:"_id"`
	IBAN         string            `json:"iban" bson:"iban"`
	Owner        string            `json:"owner" bson:"owner"`
	Value        CheckingValue     `json:"value" bson:"value"`
	LastPayments []CheckingPayment `json:"lastPayments,omitempty" bson:"lastPayments,omitempty"`
}

type CheckingValue struct {
	Amount   float32 `json:"amount" bson:"amount"`
	Currency string  `json:"currency" bson:"currency"`
}

type CheckingPayment struct {
	Id      string         `json:"id" bson:"id"`
	Type    string         `json:"type" bson:"type"`
	Value   PaymentValue   `json:"value" bson:"value"`
	Payee   Payee          `json:"payee" bson:"payee"`
	Details string         `json:"details,omitempty" bson:"details,omitempty"`
	Outcome PaymentOutcome `json:"outcome" bson:"outcome,omitempty"`
}

type InsertCheckingAccountInput struct {
	Owner string `json:"owner"`
}

type AddCheckingValueInput struct {
	Value CheckingValue `json:"value"`
}
