package model

type Order struct {
	Id           string `json:"id"`
	Dossier      string `json:"dossier"`
	Isin         string `json:"isin"`
	Type         string `json:"type"`
	Price        Price  `json:"price"`
	Quantity     int32  `json:"quantity"`
	Options      string `json:"options"`
	LeftQuantity int32  `json:"leftQuantity"`
}

type InsertOrderInput struct {
	Dossier  string `json:"dossier"`
	Isin     string `json:"isin"`
	Type     string `json:"type"`
	Price    Price  `json:"price"`
	Quantity int32  `json:"quantity"`
	Options  string `json:"options"`
}
