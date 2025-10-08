package config

import (
	"encoding/json"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type Server struct {
	Port string `json:"port" envconfig:"SERVER_PORT"`
}

type DB struct {
	Host    string `json:"host" envconfig:"DB_HOST"`
	Timeout int    `json:"timeout" envconfig:"DB_TIMEOUT"`
	Prefix  string `json:"prefix" envconfig:"DB_PREFIX"`
}

type Queue struct {
	Brokers []string `json:"brokers" envconfig:"QUEUE_BROKER"`
	Timeout int      `json:"timeout" envconfig:"QUEUE_TIMEOUT"`
	Group   string   `json:"group" envconfig:"QUEUE_GROUP"`
}

type Service struct {
	Host    string `json:"host"`
	Timeout int    `json:"timeout"`
}

func loadFromFile(path string, config any) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewDecoder(f).Decode(config)
}

func loadFromEnv(config any) error {
	return envconfig.Process("", config)
}

func LoadConfig(path string, config any) error {
	// Read from config file
	err := loadFromFile(path, config)
	if err != nil {
		return err
	}

	// Load the enviromental variables
	err = loadFromEnv(config)
	return err
}
