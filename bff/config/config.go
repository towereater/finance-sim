package config

import (
	com "mainframe-lib/common/config"
)

// Base config extension
type Config struct {
	Server   com.Server `json:"server"`
	Services struct {
		Security         com.Service `json:"security"`
		Users            com.Service `json:"users"`
		Accounts         com.Service `json:"accounts"`
		CheckingAccounts com.Service `json:"ck-accounts"`
		Dossiers         com.Service `json:"dossiers"`
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
