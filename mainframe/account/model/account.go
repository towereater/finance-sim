package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	Id    primitive.ObjectID `json:"id" bson:"_id"`
	Owner string             `json:"owner" bson:"owner"`
}

type InsertAccountInput struct {
	Owner string `json:"owner"`
}
