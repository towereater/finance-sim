package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Host string `yaml:"host" envconfig:"APP_HOST"`
		Port string `yaml:"port" envconfig:"APP_PORT"`
	} `yaml:"app"`
	Server struct {
		Users struct {
			Host string `yaml:"host" envconfig:"USERS_HOST"`
			Port string `yaml:"port" envconfig:"USERS_PORT"`
		} `yaml:"users"`
		Accounts struct {
			Host string `yaml:"host" envconfig:"ACCOUNTS_HOST"`
			Port string `yaml:"port" envconfig:"ACCOUNTS_PORT"`
		} `yaml:"accounts"`
	} `yaml:"server"`
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
