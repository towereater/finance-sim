package service

import (
	dos "mainframe-lib/dossier/model"
	xch "mainframe-lib/xchanger/model"
	"slices"
)

func ToOrder(order xch.Order) dos.Order {
	var transactions []dos.OrderTransaction
	for _, order := range slices.All(order.Transactions) {
		transactions = append(transactions, ToOrderTransaction(order))
	}

	return dos.Order{
		Id:           order.Id,
		Dossier:      order.Dossier,
		Isin:         order.Isin,
		Type:         order.Type,
		Price:        dos.Price(order.Price),
		Quantity:     order.Quantity,
		Options:      order.Options,
		LeftQuantity: order.LeftQuantity,
		Transactions: transactions,
	}
}

func ToOrderTransaction(transaction xch.OrderTransaction) dos.OrderTransaction {
	return dos.OrderTransaction{
		Dossier:   transaction.Dossier,
		Quantity:  transaction.Quantity,
		Price:     dos.Price(transaction.Price),
		Timestamp: transaction.Timestamp,
	}
}
