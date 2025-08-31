package config

import (
	com "mainframe-lib/common/config"
)

// Base config extension
type Config struct {
	com.BaseConfig
	Prefix      string `json:"prefix" envconfig:"COLL_PREFIX"`
	Collections struct {
		Users string `json:"users" envconfig:"COLL_USERS"`
	} `json:"collections"`
}

// Path and query parameters
const ContextUserId com.ContextKey = "userId"
const ContextApiKey com.ContextKey = "apiKey"
