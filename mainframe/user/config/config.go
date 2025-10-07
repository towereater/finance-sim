package config

import (
	com "mainframe-lib/common/config"
)

// Base config extension
type DBConfig struct {
	com.DBConfig
	Collections struct {
		Users string `json:"users" envconfig:"COLL_USERS"`
	} `json:"collections"`
}

type Config struct {
	Server   com.ServerConfig `json:"server"`
	DBConfig DBConfig         `json:"db"`
	Services struct {
		Security com.ServiceConfig `json:"security"`
	} `json:"services"`
}

// Path and query parameters
const ContextUserId com.ContextKey = "userId"
const ContextUsername com.ContextKey = "username"

const ContextFrom com.ContextKey = "from"
const ContextLimit com.ContextKey = "limit"
