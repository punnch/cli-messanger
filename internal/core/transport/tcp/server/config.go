package core_tcp_server

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addr string `envconfig:"ADDR" required:"true"`
}

func NewConfig() (Config, error) {
	var cfg Config

	if err := envconfig.Process("TCP", &cfg); err != nil {
		return Config{}, fmt.Errorf("process envconfig: %w", err)
	}

	return cfg, nil
}

func NewConfigMust() Config {
	cfg, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get tcp config: %w", err)
		panic(err)
	}

	return cfg
}
