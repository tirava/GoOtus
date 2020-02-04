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

func (y *ConfigYaml) readParameters() error {
	yamlFile, err := ioutil.ReadFile(y.Source)
	if err != nil {
		return fmt.Errorf("error read config file: %w", err)
	}

	if err := yaml.Unmarshal(yamlFile, &y.Config); err != nil {
		return fmt.Errorf("error unmarshal config file: %w", err)
	}

	y.Config, err = y.FillConfig(y.Config)
	if err != nil {
		return err
	}

	sort.Strings(y.NoProxyHeaders)

	return nil
}
