package config

import (
	com "mainframe-lib/common/config"
)

// Base config extension
type DB struct {
	com.DB
	Collections struct {
		Banks string `json:"banks" envconfig:"COLL_BANKS"`
		Users string `json:"users" envconfig:"COLL_USERS"`
	} `json:"collections"`
}

type Config struct {
	Server   com.Server `json:"server"`
	DBConfig DB         `json:"db"`
}

// Path and query parameters
const ContextAbi com.ContextKey = "abi"
const ContextUserId com.ContextKey = "userId"
const ContextApiKey com.ContextKey = "apiKey"
