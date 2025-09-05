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
	} `json:"collections"`
	Services struct {
		Users    string `json:"users" envconfig:"SERVICES_USERS"`
		Security string `json:"security" envconfig:"SERVICES_SECURITY"`
		Timeout  int    `json:"timeout" envconfig:"SERVICES_TIMEOUT"`
	} `json:"services"`
}

// Path and query parameters
const ContextAccount com.ContextKey = "account"
const ContextService com.ContextKey = "service"
const ContextOwner com.ContextKey = "owner"

const ContextFrom com.ContextKey = "from"
const ContextLimit com.ContextKey = "limit"
