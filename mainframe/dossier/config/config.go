package config

import (
	com "mainframe-lib/common/config"
)

// Base config extension
type DBConfig struct {
	com.DBConfig
	Collections struct {
		Dossiers string `json:"dossiers" envconfig:"COLL_DOSSIERS"`
	} `json:"collections"`
}

type Config struct {
	Server   com.ServerConfig `json:"server"`
	DBConfig DBConfig         `json:"db"`
	Services struct {
		Security         com.ServiceConfig `json:"security"`
		Users            com.ServiceConfig `json:"users"`
		Accounts         com.ServiceConfig `json:"accounts"`
		CheckingAccounts com.ServiceConfig `json:"ck-accounts"`
		Xchanger         com.ServiceConfig `json:"xchanger"`
	} `json:"services"`
}

// Path and query parameters
const ContextIsin com.ContextKey = "isin"
const ContextDossierId com.ContextKey = "dossierId"
const ContextOrderId com.ContextKey = "orderId"

const ContextFrom com.ContextKey = "from"
const ContextLimit com.ContextKey = "limit"
const ContextPage com.ContextKey = "page"
