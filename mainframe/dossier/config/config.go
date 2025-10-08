package config

import (
	com "mainframe-lib/common/config"
)

// Base config extension
type DB struct {
	com.DB
	Collections struct {
		Dossiers string `json:"dossiers" envconfig:"COLL_DOSSIERS"`
	} `json:"collections"`
}

type Config struct {
	Server   com.Server `json:"server"`
	DBConfig DB         `json:"db"`
	Services struct {
		Security         com.Service `json:"security"`
		Users            com.Service `json:"users"`
		Accounts         com.Service `json:"accounts"`
		CheckingAccounts com.Service `json:"ck-accounts"`
		Xchanger         com.Service `json:"xchanger"`
	} `json:"services"`
}

// Path and query parameters
const ContextIsin com.ContextKey = "isin"
const ContextDossierId com.ContextKey = "dossierId"
const ContextOrderId com.ContextKey = "orderId"

const ContextFrom com.ContextKey = "from"
const ContextLimit com.ContextKey = "limit"
const ContextPage com.ContextKey = "page"
