package config

import (
	com "mainframe-lib/common/config"
)

// Base config extension
type Config struct {
	com.BaseConfig
	Prefix      string `json:"prefix" envconfig:"COLL_PREFIX"`
	Collections struct {
		Dossiers string `json:"dossiers" envconfig:"COLL_DOSSIERS"`
	} `json:"collections"`
	Services struct {
		Users            string `json:"users" envconfig:"SERVICES_USERS"`
		Accounts         string `json:"accounts" envconfig:"SERVICES_ACCOUNTS"`
		CheckingAccounts string `json:"ck-accounts" envconfig:"SERVICES_CKACCOUNTS"`
		Security         string `json:"security" envconfig:"SERVICES_SECURITY"`
		Xchanger         string `json:"xchanger" envconfig:"SERVICES_XCHANGER"`
		Timeout          int    `json:"timeout" envconfig:"SERVICES_TIMEOUT"`
	} `json:"services"`
}

// Path and query parameters
const ContextDossier com.ContextKey = "dossier"

const ContextFrom com.ContextKey = "from"
const ContextLimit com.ContextKey = "limit"
