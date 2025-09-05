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
		Timeout          int    `json:"timeout" envconfig:"SERVICES_TIMEOUT"`
	} `yaml:"services"`
}

// Path and query parameters
const ContextUserId com.ContextKey = "userId"
const ContextService com.ContextKey = "service"

const ContextFrom com.ContextKey = "from"
const ContextLimit com.ContextKey = "limit"
