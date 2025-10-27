package service

import (
	dos "mainframe-lib/dossier/model"
	xch "mainframe-lib/xchanger/model"
)

func ToStock(stock xch.Stock) dos.Stock {
	return dos.Stock{
		Isin:        stock.Isin,
		Symbol:      stock.Symbol,
		Description: stock.Description,
		Type:        stock.Type,
		Prices: dos.DailyPrices{
			DailyMax:     dos.Price(stock.Prices.DailyMax),
			DailyMin:     dos.Price(stock.Prices.DailyMin),
			DailyOpening: dos.Price(stock.Prices.DailyOpening),
			DailyLast:    dos.Price(stock.Prices.DailyLast),
		},
	}
}
