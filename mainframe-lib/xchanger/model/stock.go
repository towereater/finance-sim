package model

type Stock struct {
	Isin        string `json:"isin"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Prices      struct {
		DailyMax     Price `json:"dailyMax"`
		DailyMin     Price `json:"dailyMin"`
		DailyOpening Price `json:"dailyOpening"`
	} `json:"prices"`
	SellOrders []Order `json:"sellOrders,omitempty"`
	BuyOrders  []Order `json:"buyOrders,omitempty"`
}

type InsertStockInput struct {
	Isin        string `json:"isin"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
	Type        string `json:"type"`
}
