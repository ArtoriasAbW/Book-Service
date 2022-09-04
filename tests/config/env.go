//go:build integration
// +build integration

package config

import (
	"github.com/kelseyhightower/envconfig"
)

const envPrefix = "QA"

type Config struct {
	Host       string `split_words:"true" default:"localhost:7002"`
	DbHost     string `split_words:"true" default:"localhost"`
	DbPort     string `split_words:"true" default:"5432"`
	DbUser     string `split_words:"true" default:"test_user"`
	DbPassword string `split_words:"true" default:"test_password"`
	DbName     string `split:"true" default:"book_service"`
}

func FromEnv() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process(envPrefix, cfg)
	return cfg, err
}
