package config

import (
	com "mainframe-lib/common/config"
)

// Base config extension
type Config struct {
	Server   com.ServerConfig `json:"server"`
	Services struct {
		Security         com.ServiceConfig `json:"security"`
		Users            com.ServiceConfig `json:"users"`
		Accounts         com.ServiceConfig `json:"accounts"`
		CheckingAccounts com.ServiceConfig `json:"ck-accounts"`
		Dossiers         com.ServiceConfig `json:"dossiers"`
	} `json:"services"`
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
