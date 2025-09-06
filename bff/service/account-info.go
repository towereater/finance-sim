package service

import (
	"bff/config"
	"bff/model"
	acc "mainframe-lib/account/model"
	cha "mainframe-lib/checking-account/model"
	scha "mainframe-lib/checking-account/service"
)

func GetCheckingAccountInfo(cfg config.Config, auth string, id string) (model.CheckingAccountInfo, error) {
	// Get account main details
	ckAccount, err := scha.GetAccount(cfg.Services.CheckingAccounts, cfg.Services.Timeout, auth, id)
	if err != nil {
		return model.CheckingAccountInfo{}, err
	}

	// Get latest account payments
	filter := cha.Payment{}
	filter.Payer.AccountId.Account = id
	filter.Payer.AccountId.Service = "CK"
	payments, err := scha.GetPayments(cfg.Services.CheckingAccounts, cfg.Services.Timeout, auth, filter, "", 5)
	if err != nil {
		return model.CheckingAccountInfo{}, err
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

	return ckAccountInfo, nil
}
