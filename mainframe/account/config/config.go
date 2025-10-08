package config

import (
	com "mainframe-lib/common/config"
)

// Base config extension
type DB struct {
	com.DB
	Collections struct {
		Accounts string `json:"accounts" envconfig:"COLL_ACCOUNTS"`
	} `json:"collections"`
}

type Config struct {
	Server   com.Server `json:"server"`
	DBConfig DB         `json:"db"`
	Services struct {
		Security com.Service `json:"security"`
		Users    com.Service `json:"users"`
	} `json:"services"`
}

// Path and query parameters
const ContextAccount com.ContextKey = "account"
const ContextService com.ContextKey = "service"
const ContextOwner com.ContextKey = "owner"

const ContextFrom com.ContextKey = "from"
const ContextLimit com.ContextKey = "limit"
