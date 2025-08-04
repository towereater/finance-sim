package model

type CheckingAccount struct {
	Id    string  `json:"id" bson:"_id"`
	IBAN  string  `json:"iban" bson:"iban"`
	Owner string  `json:"owner" bson:"owner"`
	Cash  float32 `json:"cash" bson:"cash"`
}

type InsertCheckingAccountInput struct {
	Owner string `json:"owner"`
}
