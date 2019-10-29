/*
 * HomeWork-9: Calendar config and logs
 * Created on 27.10.2019 15:19
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package configs implements configs helpers.
package configs

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// Config is the main config struct.
type Config struct {
	confPath string `yaml:"-"`
	DBType   string `yaml:"db_type"`
	LogLevel string `yaml:"log_level"`
}

// NewConfig creates new config struct.
func NewConfig(confPath string) *Config {
	return &Config{confPath: confPath}
}

// ReadParameters reads config from file.
func (c *Config) ReadParameters() error {
	yamlFile, err := ioutil.ReadFile(c.confPath)
	if err != nil {
		return fmt.Errorf("error read config file: %w", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return fmt.Errorf("error unmarshal config file: %w", err)
	}
	if c.DBType == "" {
		c.DBType = "map"
	}
	if c.LogLevel == "" {
		c.LogLevel = "info"
	}
	return nil
}
