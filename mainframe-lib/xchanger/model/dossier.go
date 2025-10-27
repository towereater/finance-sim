package model

type Dossier struct {
	Id         string         `json:"id"`
	Name       string         `json:"name"`
	Surname    string         `json:"surname"`
	Birth      string         `json:"birth"`
	Abi        string         `json:"abi"`
	ExternalId string         `json:"externalId"`
	IBAN       string         `json:"iban"`
	Stocks     []DossierStock `json:"stocks,omitempty"`
	Value      DossierValue   `json:"value"`
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
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Birth      string `json:"birth"`
	ExternalId string `json:"externalId"`
	IBAN       string `json:"iban"`
}
