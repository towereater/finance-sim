package config

import (
	com "mainframe-lib/common/config"
)

// Base config extension
type DB struct {
	com.DB
	Collections struct {
		Accounts string `json:"accounts" envconfig:"COLL_ACCOUNTS"`
		Payments string `json:"payments" envconfig:"COLL_PAYMENTS"`
	} `json:"collections"`
}

type Queue struct {
	com.Queue
	Topics struct {
		Payments string `json:"payments" envconfig:"TOPIC_PAYMENTS"`
	} `json:"topics"`
}

type Config struct {
	Server   com.Server `json:"server"`
	DB       DB         `json:"db"`
	Queue    Queue      `json:"queue"`
	Services struct {
		Security com.Service `json:"security"`
		Users    com.Service `json:"users"`
		Accounts com.Service `json:"accounts"`
	} `json:"services"`
}

// Path and query parameters
const ContextAccountId com.ContextKey = "accountId"
const ContextPaymentId com.ContextKey = "paymentId"

const ContextFrom com.ContextKey = "from"
const ContextLimit com.ContextKey = "limit"
