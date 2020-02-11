// Package configs implements config interface.
package configs

import (
	"fmt"
	"strings"

	"gitlab.com/tirava/image-previewer/internal/configs/inmemory"
	"gitlab.com/tirava/image-previewer/internal/configs/yaml"
	"gitlab.com/tirava/image-previewer/internal/models"
)

// NewConfig returns config implementer.
func NewConfig(implementer string) (models.Configer, error) {
	switch {
	case implementer == models.InMemory:
		return inmemory.NewConfig(implementer)
	case strings.Contains(implementer, ".yml") || strings.Contains(implementer, ".yaml"):
		return yaml.NewConfig(implementer)
	}

	return nil, fmt.Errorf("incorrect config implementer name: %s", implementer)
}
