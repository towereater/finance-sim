package service

import (
	"bff/config"
	"bff/model"
	"fmt"
	acc "mainframe-lib/account/model"
	scha "mainframe-lib/checking-account/service"
	sdos "mainframe-lib/dossier/service"
	"net/http"
)

func GetCheckingAccountInfo(cfg config.Config, auth string, userId string, accountId string) (model.CheckingAccountInfo, int, error) {
	// Get account main details
	ckAccount, status, err := scha.GetAccount(cfg.Services.CheckingAccounts, auth, accountId)
	if err != nil {
		return model.CheckingAccountInfo{}, status, err
	}
	if status != http.StatusOK {
		return model.CheckingAccountInfo{}, status, fmt.Errorf("account not found")
	}

	// Construct account info
	ckAccountInfo := model.CheckingAccountInfo{
		AccountInfo: model.AccountInfo{
			AccountId: acc.AccountId{
				Account: accountId,
				Service: "CK",
			},
		},
		IBAN:         ckAccount.IBAN,
		Value:        ckAccount.Value,
		LastPayments: ckAccount.LastPayments,
	}

	return ckAccountInfo, status, nil
}

func GetDossierInfo(cfg config.Config, auth string, userId string, dossierId string) (model.DossierInfo, int, error) {
	// Get account main details
	dossier, status, err := sdos.GetDossier(cfg.Services.Dossiers, auth, dossierId)
	if err != nil {
		return model.DossierInfo{}, status, err
	}
	if status != http.StatusOK {
		return model.DossierInfo{}, status, fmt.Errorf("account not found")
	}

	// Construct account info
	dossierInfo := model.DossierInfo{
		AccountInfo: model.AccountInfo{
			AccountId: acc.AccountId{
				Account: dossierId,
				Service: "DS",
			},
		},
		Value: dossier.Value,
	}

	return dossierInfo, status, nil
}
