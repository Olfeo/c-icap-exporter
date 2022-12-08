package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServicePort    string `env:"SERVICE_PORT" envDefault:"8080"`
	ICAPPort       string `env:"ICAP_PORT" envDefault:"1344"`
	ICAPAddress    string `env:"ICAP_ADDRESS" envDefault:"localhost"`
	ICAPClientPath string `env:"ICAP_CLIENT_PATH" envDefault:"/usr/local/c-icap/bin/c-icap-client"`
}

func NewConfig() (*Config, error) {
	config := new(Config)
	if err := env.Parse(config); err != nil {
		return nil, fmt.Errorf("env.Parse error: %w", err)
	}
	return config, nil
}
