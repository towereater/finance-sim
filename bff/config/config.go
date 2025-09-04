package config

import (
	com "mainframe-lib/common/config"
)

type Config struct {
	com.BaseConfig
	Services struct {
		Security string `json:"security" envconfig:"SERVICES_SECURITY"`
		Users    string `json:"users" envconfig:"SERVICES_USERS"`
		//Accounts string `yaml:"accounts" envconfig:"SERVICES_ACCOUNTS"`
		Timeout int `json:"timeout" envconfig:"SERVICES_TIMEOUT"`
	} `yaml:"services"`
}

// Path and query parameters
const ContextUserId com.ContextKey = "userId"
