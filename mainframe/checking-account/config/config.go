package config

import (
	com "mainframe-lib/common/config"
)

// Base config extension
type Config struct {
	com.BaseConfig
	Prefix      string `json:"prefix" envconfig:"COLL_PREFIX"`
	Collections struct {
		Accounts string `json:"accounts" envconfig:"COLL_ACCOUNTS"`
		Payments string `json:"payments" envconfig:"COLL_PAYMENTS"`
	} `json:"collections"`
	Services struct {
		Users    string `json:"users" envconfig:"SERVICES_USERS"`
		Accounts string `json:"accounts" envconfig:"SERVICES_ACCOUNTS"`
		Security string `json:"security" envconfig:"SERVICES_SECURITY"`
		Timeout  int    `json:"timeout" envconfig:"SERVICES_TIMEOUT"`
	} `json:"services"`
}

// Path and query parameters
const ContextAccountId com.ContextKey = "accountId"
const ContextPaymentId com.ContextKey = "paymentId"

const ContextFrom com.ContextKey = "from"
const ContextLimit com.ContextKey = "limit"
