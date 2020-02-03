/*
 * Project: Image Previewer
 * Created on 20.01.2020 11:44
 * Copyright (c) 2020 - Eugene Klimov
 */

package yaml

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	"gitlab.com/tirava/image-previewer/internal/models"

	"gopkg.in/yaml.v2"
)

// ConfigYaml is yaml config.
type ConfigYaml struct {
	models.Config
}

// NewConfig returns new config struct.
func NewConfig(confPath string) (ConfigYaml, error) {
	y := ConfigYaml{
		models.Config{Source: confPath},
	}

	return y, y.readParameters()
}

// GetConfig got yaml config struct.
func (y ConfigYaml) GetConfig() models.Config {
	return y.Config
}

// SetConfig sets new yaml config struct.
func (ConfigYaml) SetConfig(conf models.Config) {
}

// nolint )
func (y *ConfigYaml) readParameters() error {
	yamlFile, err := ioutil.ReadFile(y.Source)
	if err != nil {
		return fmt.Errorf("error read config file: %w", err)
	}

	err = yaml.Unmarshal(yamlFile, &y.Config)

	if err != nil {
		return fmt.Errorf("error unmarshal config file: %w", err)
	}

	defConfig := y.GetDefaults()

	if y.Logger == "" {
		y.Logger = defConfig.Logger
	}
	if y.LogFile == "" {
		y.LogFile = defConfig.LogFile
	}
	if y.LogLevel == "" {
		y.LogLevel = defConfig.LogLevel
	}
	if y.ListenHTTP == "" {
		y.ListenHTTP = defConfig.ListenHTTP
	}
	if y.ListenPrometheus == "" {
		y.ListenPrometheus = defConfig.ListenPrometheus
	}
	if y.ListenPprof == "" {
		y.ListenPprof = defConfig.ListenPprof
	}
	if y.Previewer == "" {
		y.Previewer = defConfig.Previewer
	}
	if y.ImageURLEncoder == "" {
		y.ImageURLEncoder = defConfig.ImageURLEncoder
	}
	if y.Cacher == "" {
		y.Cacher = defConfig.Cacher
	}
	if y.Storager == "" {
		y.Storager = defConfig.Storager
	}
	if y.StoragePath == "" {
		y.StoragePath = defConfig.StoragePath
	}
	if y.MaxCacheItems == 0 {
		y.MaxCacheItems = defConfig.MaxCacheItems
	}

	if y.ImageURLEncoder == "" {
		y.ImageURLEncoder = defConfig.ImageURLEncoder
	}

	if y.Cacher == "" {
		y.Cacher = defConfig.Cacher
	}

	if y.Storager == "" {
		y.Storager = defConfig.Storager
	}

	if y.StoragePath == "" {
		y.StoragePath = defConfig.StoragePath
	}

	for i, s := range y.NoProxyHeaders {
		y.NoProxyHeaders[i] = strings.ToLower(s)
	}

	sort.Strings(y.NoProxyHeaders)

	return nil
}
