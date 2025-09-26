package config

import (
	com "mainframe-lib/common/config"
)

type Config struct {
	com.BaseConfig
	Services struct {
		Security         string `json:"security" envconfig:"SERVICES_SECURITY"`
		Users            string `json:"users" envconfig:"SERVICES_USERS"`
		Accounts         string `json:"accounts" envconfig:"SERVICES_ACCOUNTS"`
		CheckingAccounts string `json:"checking-accounts" envconfig:"SERVICES_CKACCOUNTS"`
		Dossiers         string `json:"dossiers" envconfig:"SERVICES_DOSSIERS"`
		Timeout          int    `json:"timeout" envconfig:"SERVICES_TIMEOUT"`
	} `yaml:"services"`
}

// Path and query parameters
const ContextUserId com.ContextKey = "userId"
const ContextService com.ContextKey = "service"
const ContextPaymentId com.ContextKey = "paymentId"
const ContextIsin com.ContextKey = "isin"
const ContextOrderId com.ContextKey = "orderId"

const ContextFrom com.ContextKey = "from"
const ContextLimit com.ContextKey = "limit"
const ContextPage com.ContextKey = "page"
