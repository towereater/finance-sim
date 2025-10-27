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

type DossierDto struct {
	Id              string         `json:"id"`
	Owner           string         `json:"owner"`
	CheckingAccount acc.AccountId  `json:"checkingAccount"`
	XChangerDossier string         `json:"xchangerDossier"`
	Stocks          []DossierStock `json:"stocks,omitempty"`
	Value           DossierValue   `json:"value"`
}

type DossierStock struct {
	Isin      string `json:"isin"`
	Total     int32  `json:"total"`
	Available int32  `json:"available"`
}

type DossierValue struct {
	Value     Price  `json:"value"`
	Timestamp string `json:"timestamp"`
}

type InsertDossierInput struct {
	Owner           string        `json:"owner"`
	CheckingAccount acc.AccountId `json:"checkingAccount"`
}
