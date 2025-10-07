package config

import (
	com "mainframe-lib/common/config"
)

// Base config extension
type DBConfig struct {
	com.DBConfig
	Collections struct {
		Accounts string `json:"accounts" envconfig:"COLL_ACCOUNTS"`
		Payments string `json:"payments" envconfig:"COLL_PAYMENTS"`
	} `json:"collections"`
}

type Config struct {
	Server   com.ServerConfig `json:"server"`
	DBConfig DBConfig         `json:"db"`
	Services struct {
		Security com.ServiceConfig `json:"security"`
		Users    com.ServiceConfig `json:"users"`
		Accounts com.ServiceConfig `json:"accounts"`
	} `json:"services"`
}

// Path and query parameters
const ContextAccountId com.ContextKey = "accountId"
const ContextPaymentId com.ContextKey = "paymentId"

const ContextFrom com.ContextKey = "from"
const ContextLimit com.ContextKey = "limit"
