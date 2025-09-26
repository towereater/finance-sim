package service

import (
	dos "mainframe-lib/dossier/model"
	xch "mainframe-lib/xchanger/model"
)

func ToOrder(order xch.Order) dos.Order {
	return dos.Order{
		Id:           order.Id,
		Dossier:      order.Dossier,
		Isin:         order.Isin,
		Type:         order.Type,
		Price:        dos.Price(order.Price),
		Quantity:     order.Quantity,
		Options:      order.Options,
		LeftQuantity: order.LeftQuantity,
	}
}
