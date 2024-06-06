package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host string `yaml:"host" envconfig:"SERVER_HOST"`
		Port string `yaml:"port" envconfig:"SERVER_PORT"`
	} `yaml:"server"`
	DB struct {
		Host    string `yaml:"host" envconfig:"DB_HOST"`
		Port    string `yaml:"port" envconfig:"DB_PORT"`
		Timeout int    `yaml:"timeout" envconfig:"DB_TIMEOUT"`
	} `yaml:"db"`
}

var AppConfig Config

func readConfig(path string) error {
	// Read entire config file
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// Conversion of the yaml to struct
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		return err
	}

	return nil
}

func readEnv() error {
	// Loads the enviromental variables
	err := envconfig.Process("", &AppConfig)
	if err != nil {
		return err
	}

	return nil
}

func LoadConfig(path string) error {
	// Reading config file
	err := readConfig(path)
	if err != nil {
		return err
	}

	// Setting environmental variables
	err = readEnv()
	if err != nil {
		return err
	}

	return nil
}
