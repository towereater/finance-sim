package config

import (
	com "mainframe-lib/common/config"
)

// Base config extension
type Config struct {
	com.BaseConfig
	Prefix      string `json:"prefix" envconfig:"COLL_PREFIX"`
	Collections struct {
		Users string `yaml:"users" envconfig:"COLL_USERS"`
	} `yaml:"collections"`
	Services struct {
		Security string `yaml:"security" envconfig:"SERVICES_SECURITY"`
		Timeout  int    `yaml:"timeout" envconfig:"SERVICES_TIMEOUT"`
	} `yaml:"services"`
}

// Path and query parameters
const ContextUserId com.ContextKey = "userId"
const ContextUsername com.ContextKey = "username"

const ContextFrom com.ContextKey = "from"
const ContextLimit com.ContextKey = "limit"
