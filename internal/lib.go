package config

import "github.com/kelseyhightower/envconfig"

func LoadConfig() (Config, error) {
	var config Config
	err := envconfig.Process("gouh", &config)
	return config, err
}

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Address string `default:"0.0.0.0"`
	Port    string `default:"8080"`
}
