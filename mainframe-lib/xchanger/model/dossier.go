package model

type Dossier struct {
	Id string `json:"id"`
}

type InsertDossierInput struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Birth      string `json:"birth"`
	ExternalId string `json:"externalId"`
	IBAN       string `json:"iban"`
}
