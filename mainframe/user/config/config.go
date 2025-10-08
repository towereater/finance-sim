package config

import (
	com "mainframe-lib/common/config"
)

// Base config extension
type DB struct {
	com.DB
	Collections struct {
		Users string `json:"users" envconfig:"COLL_USERS"`
	} `json:"collections"`
}

type Config struct {
	Server   com.Server `json:"server"`
	DBConfig DB         `json:"db"`
	Services struct {
		Security com.Service `json:"security"`
	} `json:"services"`
}

// Path and query parameters
const ContextUserId com.ContextKey = "userId"
const ContextUsername com.ContextKey = "username"

const ContextFrom com.ContextKey = "from"
const ContextLimit com.ContextKey = "limit"
