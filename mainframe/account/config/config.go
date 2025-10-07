package config

import (
	com "mainframe-lib/common/config"
)

// Base config extension
type DBConfig struct {
	com.DBConfig
	Collections struct {
		Accounts string `json:"accounts" envconfig:"COLL_ACCOUNTS"`
	} `json:"collections"`
}

type Config struct {
	Server   com.ServerConfig `json:"server"`
	DBConfig DBConfig         `json:"db"`
	Services struct {
		Security com.ServiceConfig `json:"security"`
		Users    com.ServiceConfig `json:"users"`
	} `json:"services"`
}

// Path and query parameters
const ContextAccount com.ContextKey = "account"
const ContextService com.ContextKey = "service"
const ContextOwner com.ContextKey = "owner"

const ContextFrom com.ContextKey = "from"
const ContextLimit com.ContextKey = "limit"
