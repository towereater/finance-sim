package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	Id    primitive.ObjectID `json:"id" bson:"_id"`
	IBAN  string             `json:"iban" bson:"iban"`
	Owner string             `json:"owner" bson:"owner"`
	Cash  float32            `json:"cash" bson:"cash"`
}
