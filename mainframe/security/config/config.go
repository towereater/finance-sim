package config

import (
	lib "mainframe-lib/common/config"
)

// Base config extension
type Config struct {
	lib.BaseConfig
	Prefix      string `json:"prefix" envconfig:"COLL_PREFIX"`
	Collections struct {
		Users string `json:"users" envconfig:"COLL_USERS"`
	} `json:"collections"`
}

// Path and query parameters
const ContextUserId lib.ContextKey = "userId"
const ContextApiKey lib.ContextKey = "apiKey"
