package model

type AccountId struct {
	Account string `json:"account" bson:"account"`
	Service string `json:"service" bson:"service"`
}

type InsertAccountInput struct {
	Id    AccountId `json:"id"`
	Owner string    `json:"owner"`
}
