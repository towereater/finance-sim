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
	DB    DB    `json:"db"`
	Queue Queue `json:"queue"`
}
