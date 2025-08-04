package model

type Dossier struct {
	Id              string    `json:"id" bson:"_id"`
	Owner           string    `json:"owner" bson:"owner"`
	CheckingAccount AccountId `json:"checkingAccount" bson:"checkingAccount"`
	XChangerDossier string    `json:"xchangerDossier" bson:"xchangerDossier"`
}

type InsertDossierInput struct {
	Owner           string    `json:"owner"`
	CheckingAccount AccountId `json:"checkingAccount"`
}
