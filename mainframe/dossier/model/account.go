package model

type Account struct {
	Id    AccountId `json:"id" bson:"_id"`
	Owner string    `json:"owner" bson:"owner"`
}

type AccountId struct {
	Account string `json:"account" bson:"account"`
	Service string `json:"service" bson:"service"`
}

type InsertAccountInput struct {
	Id    AccountId `json:"id"`
	Owner string    `json:"owner"`
}
