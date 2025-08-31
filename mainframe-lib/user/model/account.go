package model

type Account struct {
	Id AccountId `json:"id" bson:"id"`
}

type AccountId struct {
	Account string `json:"account" bson:"account"`
	Service string `json:"service" bson:"service"`
}

type InsertAccountInput struct {
	Id AccountId `json:"id"`
}

type DeleteAccountInput struct {
	Id AccountId `json:"id"`
}
