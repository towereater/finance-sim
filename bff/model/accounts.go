package model

import (
	acc "mainframe-lib/account/model"
	cha "mainframe-lib/checking-account/model"
	dos "mainframe-lib/dossier/model"
)

type AccountInfo struct {
	AccountId acc.AccountId `json:"id"`
}

type CheckingAccountInfo struct {
	AccountInfo
	IBAN         string                `json:"iban"`
	Value        cha.CheckingValue     `json:"value"`
	LastPayments []cha.CheckingPayment `json:"lastPayments,omitempty"`
}

type DossierInfo struct {
	AccountInfo
	Value dos.DossierValue `json:"value"`
}

type GetAccountsOutput struct {
	Accounts []any `json:"accounts"`
}
