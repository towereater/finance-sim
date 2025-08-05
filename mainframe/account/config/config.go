package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type DBConfig struct {
	Host    string `yaml:"host" envconfig:"DB_HOST"`
	Timeout int    `yaml:"timeout" envconfig:"DB_TIMEOUT"`
}

type Config struct {
	Server struct {
		Port string `yaml:"port" envconfig:"SERVER_PORT"`
	} `yaml:"server"`
	DB          DBConfig `yaml:"db"`
	Prefix      string   `yaml:"prefix" envconfig:"COLL_PREFIX"`
	Collections struct {
		Accounts string `yaml:"accounts" envconfig:"COLL_ACCOUNTS"`
	} `yaml:"collections"`
	Services struct {
		Users    string `yaml:"users" envconfig:"SERVICES_USERS"`
		Security string `yaml:"security" envconfig:"SERVICES_SECURITY"`
		Timeout  int    `yaml:"timeout" envconfig:"SERVICES_TIMEOUT"`
	} `yaml:"services"`
}

func readConfig(path string) (Config, error) {
	//Read entire config file
	f, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer f.Close()

	// Conversion of the yaml to struct
	var config Config

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	return config, err
}

func readEnv(config Config) (Config, error) {
	// Loads the enviromental variables
	err := envconfig.Process("", &config)
	return config, err
}

func LoadConfig(path string) (Config, error) {
	// Reading config file
	config, err := readConfig(path)
	if err != nil {
		return Config{}, err
	}

	// Setting environmental variables
	config, err = readEnv(config)

	return config, err
}
