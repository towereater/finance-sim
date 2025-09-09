package model

import (
	acc "mainframe-lib/account/model"
)

type Dossier struct {
	Id              string        `json:"id" bson:"_id"`
	Owner           string        `json:"owner" bson:"owner"`
	CheckingAccount acc.AccountId `json:"checkingAccount" bson:"checkingAccount"`
	XChangerDossier string        `json:"xchangerDossier" bson:"xchangerDossier"`
}

type InsertDossierInput struct {
	Owner           string        `json:"owner"`
	CheckingAccount acc.AccountId `json:"checkingAccount"`
}
