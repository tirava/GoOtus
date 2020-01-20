/*
 * Project: Image Previewer
 * Created on 20.01.2020 11:06
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package configs implements config interface.
package configs

import (
	"errors"
	"strings"

	"gitlab.com/tirava/image-previewer/internal/models"

	"gitlab.com/tirava/image-previewer/internal/configs/inmemory"
	"gitlab.com/tirava/image-previewer/internal/configs/yaml"
)

// Configer is the main interface for configs.
type Configer interface {
	GetConfig() models.Config
	SetConfig(config models.Config) error
}

// NewConfig returns config implementer.
func NewConfig(implementer string) (Configer, error) {
	switch {
	case implementer == "inmemory":
		return inmemory.NewConfig(implementer)
	case strings.Contains(implementer, ".yml") || strings.Contains(implementer, ".yaml"):
		return yaml.NewConfig(implementer)
	}

	return nil, errors.New("incorrect config implementer name")
}
