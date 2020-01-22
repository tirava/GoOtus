/*
 * Project: Image Previewer
 * Created on 20.01.2020 11:44
 * Copyright (c) 2020 - Eugene Klimov
 */

package inmemory

import (
	"gitlab.com/tirava/image-previewer/internal/models"
)

// ConfigInMemory is map config.
type ConfigInMemory struct {
	models.Config
}

// NewConfig returns new config struct.
func NewConfig(config string) (ConfigInMemory, error) {
	m := ConfigInMemory{
		models.Config{}.GetDefaults(),
	}
	m.Source = config
	m.LogFile = "/dev/stderr"

	return m, nil
}

// GetConfig got config struct.
func (m ConfigInMemory) GetConfig() models.Config {
	return m.Config
}

// SetConfig sets config struct.
func (ConfigInMemory) SetConfig(conf models.Config) {
}
