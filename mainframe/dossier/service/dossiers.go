package service

import (
	dos "mainframe-lib/dossier/model"
	xch "mainframe-lib/xchanger/model"
	"slices"
)

func ToDossierDto(dossier dos.Dossier, xchDossier xch.Dossier) dos.DossierDto {
	var stocks []dos.DossierStock
	for _, stock := range slices.All(xchDossier.Stocks) {
		stocks = append(stocks, ToDossierStock(stock))
	}

	return dos.DossierDto{
		Id:              dossier.Id,
		Owner:           dossier.Owner,
		CheckingAccount: dossier.CheckingAccount,
		XChangerDossier: dossier.XChangerDossier,
		Stocks:          stocks,
		Value:           ToDossierValue(xchDossier.Value),
	}
}

func ToDossierStock(stock xch.DossierStock) dos.DossierStock {
	return dos.DossierStock{
		Isin:      stock.Isin,
		Total:     stock.Total,
		Available: stock.Available,
	}
}

func ToDossierValue(value xch.DossierValue) dos.DossierValue {
	return dos.DossierValue{
		Value:     dos.Price(value.Value),
		Timestamp: value.Timestamp,
	}
}
