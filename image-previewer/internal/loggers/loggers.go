// Package loggers implements logger interface.
package loggers

import (
	"errors"
	"io"

	"gitlab.com/tirava/image-previewer/internal/loggers/logrus"
	"gitlab.com/tirava/image-previewer/internal/loggers/logstd"
	"gitlab.com/tirava/image-previewer/internal/models"
)

// NewLogger returns logger implementer.
func NewLogger(implementer, level string, output io.Writer) (models.Loggerer, error) {
	switch implementer {
	case models.LogRus:
		return logrus.NewLogger(level, output)
	case models.LogStd:
		return logstd.NewLogger(level, output)
	}

	return nil, errors.New("incorrect logger implementer name")
}
