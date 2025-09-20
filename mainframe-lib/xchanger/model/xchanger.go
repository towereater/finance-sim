package model

type XChangerDossier struct {
	Id string `json:"id"`
}

type InsertXChangerDossierInput struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Birth      string `json:"birth"`
	ExternalId string `json:"externalId"`
	IBAN       string `json:"iban"`
}
