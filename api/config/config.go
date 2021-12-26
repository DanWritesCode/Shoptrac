package config

import (
	"encoding/json"
	"io/ioutil"
)

// DBConfig holds the database options.
type DBConfig struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Address  string `json:"address,omitempty"`
	Database string `json:"database,omitempty"`
}

type AppConfig struct {
	ClientId     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
  RedirectURL string `json:"redirect_url,omitempty"`
}

// Config holds configuration data for the API.
type Config struct {
	Host    string     `json:"host,omitempty"`
	Port    int        `json:"port,omitempty"`
	Debug   bool       `json:"debug,omitempty"`
	Timeout int        `json:"timeout,omitempty"`
	LogFile string     `json:"log_file,omitempty"`
	DB      *DBConfig  `json:"db,omitempty"`
	App     *AppConfig `json:"app,omitempty"`
}

var cfg *Config

// NewConfig returns a new config from a given path.
func NewConfig(path string) (*Config, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = json.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	cfg = config

	return config, nil
}

// GetConfig returns the global config.
func GetConfig() *Config {
	return cfg
}
