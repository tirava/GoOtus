/*
 * Project: Image Previewer
 * Created on 20.01.2020 11:06
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package models implements models.
package models

import "gitlab.com/tirava/image-previewer/internal/domain/entities"

const maxCacheItems = 128

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
	ListenPprof      string                 `yaml:"pprof_listen"`
	Previewer        string                 `yaml:"previewer"`
	Interpolation    entities.Interpolation `yaml:"interpolation"`
	NoProxyHeaders   []string               `yaml:"no_proxy_headers"`
	ImageURLEncoder  string                 `yaml:"img_url_encoder"`
	Cacher           string                 `yaml:"cache"`
	MaxCacheItems    int                    `yaml:"max_cache_items"`
	Storager         string                 `yaml:"storage"`
	StoragePath      string                 `yaml:"storage_path"`
}

// GetDefaults returns default config fields.
func (c Config) GetDefaults() Config {
	return Config{
		Logger:           "logstd",
		LogFile:          "previewer.log",
		LogLevel:         "info",
		ListenHTTP:       ":8080",
		ListenPrometheus: ":9180",
		ListenPprof:      ":8181",
		Previewer:        "xdraw",
		ImageURLEncoder:  "md5",
		Cacher:           "lru",
		MaxCacheItems:    maxCacheItems,
		Storager:         "disk",
		StoragePath:      "cache",
	}
}
