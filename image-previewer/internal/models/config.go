/*
 * Project: Image Previewer
 * Created on 20.01.2020 11:06
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package models implements config model.
package models

import "gitlab.com/tirava/image-previewer/internal/domain/entities"

// Config is the base config struct.
type Config struct {
	Source        string                 `yaml:"-"`
	LogFile       string                 `yaml:"log_file"`
	LogLevel      string                 `yaml:"log_level"`
	ListenHTTP    string                 `yaml:"http_listen"`
	Previewer     string                 `yaml:"previewer"`
	Interpolation entities.Interpolation `yaml:"interpolation"`
}

// GetDefaults returns default config fields.
func (c Config) GetDefaults() Config {
	return Config{
		LogFile:    "previewer.log",
		LogLevel:   "info",
		ListenHTTP: "8080",
		Previewer:  "xdraw",
	}
}
