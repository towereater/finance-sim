package model

type ApiKey struct {
	Id  string `json:"id" bson:"_id"`
	Abi string `json:"abi" bson:"abi"`
	Cab string `json:"cab" bson:"cab"`
}

type InsertApiKeyInput struct {
	Abi string `json:"abi" bson:"abi"`
	Cab string `json:"cab" bson:"cab"`
}
