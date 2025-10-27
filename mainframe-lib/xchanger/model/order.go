package model

type Order struct {
	Id           string             `json:"id"`
	Dossier      string             `json:"dossier"`
	Isin         string             `json:"isin"`
	Type         string             `json:"type"`
	Price        Price              `json:"price"`
	Quantity     int32              `json:"quantity"`
	Options      string             `json:"options"`
	LeftQuantity int32              `json:"leftQuantity"`
	Transactions []OrderTransaction `json:"transactions,omitempty"`
}

type OrderTransaction struct {
	Dossier   string `json:"dossier"`
	Quantity  int32  `json:"quantity"`
	Price     Price  `json:"price"`
	Timestamp string `json:"timestamp"`
}

type InsertOrderInput struct {
	Dossier  string `json:"dossier"`
	Isin     string `json:"isin"`
	Type     string `json:"type"`
	Price    Price  `json:"price"`
	Quantity int32  `json:"quantity"`
	Options  string `json:"options"`
}
