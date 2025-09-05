package model

import (
	acc "mainframe-lib/account/model"
	cha "mainframe-lib/checking-account/model"
)

type AccountInfo struct {
	AccountId acc.AccountId `json:"id"`
}

type CheckingAccountInfo struct {
	AccountInfo
	IBAN  string            `json:"iban"`
	Value cha.CheckingValue `json:"value" bson:"value"`
}

type GetAccountsOutput struct {
	Accounts []any `json:"accounts"`
}
