package model

type Bank struct {
	Abi            string `json:"abi" bson:"_id"`
	XchangerApiKey string `json:"xchangerApiKey,omitempty" bson:"xchangerApiKey,omitempty"`
}

type InsertBankInput struct {
	Abi            string `json:"abi"`
	XchangerApiKey string `json:"xchangerApiKey,omitempty"`
}
