// Package config implements config working.
package config

import env "github.com/caarlos0/env/v6"

// DBConfig base struct
type DBConfig struct {
	DBHost     string `env:"DB_HOST" envDefault:"localhost"`
	DBPort     string `env:"DB_PORT" envDefault:"5432"`
	DBName     string `env:"DB_NAME"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASS"`
}

// NewDBConfig returns new config.
func NewDBConfig() (*DBConfig, error) {
	cfg := &DBConfig{}
	if err := env.Parse(cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
