/*
 * Project: Image Previewer
 * Created on 20.01.2020 11:06
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package models implements models.
package models

import "gitlab.com/tirava/image-previewer/internal/domain/entities"

// Configer is the main interface for configs.
type Configer interface {
	GetConfig() Config
	SetConfig(Config)
}

// Config is the base config struct.
type Config struct {
	Source           string                 `yaml:"-"`
	Logger           string                 `yaml:"logger"`
	LogFile          string                 `yaml:"log_file"`
	LogLevel         string                 `yaml:"log_level"`
	ListenHTTP       string                 `yaml:"http_listen"`
	ListenPrometheus string                 `yaml:"prometheus_listen"`
	Previewer        string                 `yaml:"previewer"`
	Interpolation    entities.Interpolation `yaml:"interpolation"`
	NoProxyHeaders   []string               `yaml:"no_proxy_headers"`
}

// GetDefaults returns default config fields.
func (c Config) GetDefaults() Config {
	return Config{
		Logger:           "logstd",
		LogFile:          "previewer.log",
		LogLevel:         "info",
		ListenHTTP:       ":8080",
		ListenPrometheus: ":9180",
		Previewer:        "xdraw",
	}
}
