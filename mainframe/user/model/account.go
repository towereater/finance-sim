package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	Id      primitive.ObjectID `json:"id" bson:"id"`
	Service string             `json:"service" bson:"service"`
}

type InsertAccountInput struct {
	Id      primitive.ObjectID `json:"id"`
	Service string             `json:"service"`
}
