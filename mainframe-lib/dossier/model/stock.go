package model

type Stock struct {
	Isin        string      `json:"isin"`
	Symbol      string      `json:"symbol"`
	Description string      `json:"description"`
	Type        string      `json:"type"`
	Prices      DailyPrices `json:"prices"`
}

type DailyPrices struct {
	DailyMax     Price `json:"dailyMax"`
	DailyMin     Price `json:"dailyMin"`
	DailyOpening Price `json:"dailyOpening"`
	DailyLast    Price `json:"dailyLast"`
}
