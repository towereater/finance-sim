package service

import (
	dos "mainframe-lib/dossier/model"
	xch "mainframe-lib/xchanger/model"
	"slices"
)

func ToStock(stock xch.Stock) dos.Stock {
	var sellOrders []dos.Order
	for _, order := range slices.All(stock.SellOrders) {
		sellOrders = append(sellOrders, ToOrder(order))
	}

	var buyOrders []dos.Order
	for _, order := range slices.All(stock.BuyOrders) {
		buyOrders = append(buyOrders, ToOrder(order))
	}

	return dos.Stock{
		Isin:        stock.Isin,
		Symbol:      stock.Symbol,
		Description: stock.Description,
		Type:        stock.Type,
		Prices: dos.DailyPrices{
			DailyMax:     dos.Price(stock.Prices.DailyMax),
			DailyMin:     dos.Price(stock.Prices.DailyMin),
			DailyOpening: dos.Price(stock.Prices.DailyOpening),
		},
		SellOrders: sellOrders,
		BuyOrders:  buyOrders,
	}
}
