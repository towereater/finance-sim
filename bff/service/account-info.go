package service

import (
	"bff/config"
	"bff/model"
	acc "mainframe-lib/account/model"
	cha "mainframe-lib/checking-account/model"
	scha "mainframe-lib/checking-account/service"
)

func GetCheckingAccountInfo(cfg config.Config, auth string, id string) (model.CheckingAccountInfo, int, error) {
	// Get account main details
	ckAccount, status, err := scha.GetAccount(cfg.Services.CheckingAccounts, auth, id)
	if err != nil {
		return model.CheckingAccountInfo{}, status, err
	}

	// Get latest account payments
	filter := cha.Payment{}
	filter.Payer.AccountIdentification.Type = "ID"
	filter.Payer.AccountIdentification.Value = id
	payments, status, err := scha.GetPayments(cfg.Services.CheckingAccounts, auth, filter, "", 5)
	if err != nil {
		return model.CheckingAccountInfo{}, status, err
	}

	// Construct account info
	ckAccountInfo := model.CheckingAccountInfo{
		AccountInfo: model.AccountInfo{
			AccountId: acc.AccountId{
				Account: id,
				Service: "CK",
			},
		},
		IBAN:  ckAccount.IBAN,
		Value: ckAccount.Value,
	}
	if len(payments) > 0 {
		ckAccountInfo.LatestPayments = payments
	}

	return ckAccountInfo, status, nil
}
