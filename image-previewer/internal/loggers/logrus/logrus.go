// Package logrus implements logrus logger.
package logrus

import (
	"io"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gitlab.com/tirava/image-previewer/internal/models"
)

// Logger is the base struct for logrus logger.
type Logger struct {
	logger *logrus.Logger
	models.Logger
}

// NewLogger returns new logger.
func NewLogger(level string, output io.Writer) (Logger, error) {
	lg := Logger{
		logger: logrus.New(),
	}

	if level == "none" {
		lg.logger.SetOutput(ioutil.Discard)
	} else {
		lg.logger.SetOutput(output)
	}

	switch level {
	case "debug":
		lg.logger.SetLevel(logrus.DebugLevel)
	case "info":
		lg.logger.SetLevel(logrus.InfoLevel)
	case "warn":
		lg.logger.SetLevel(logrus.WarnLevel)
	default:
		lg.logger.SetLevel(logrus.ErrorLevel)
	}

	return lg, nil
}

// GetLogger returns global logger.
func (l Logger) GetLogger() models.Loggerer {
	return l
}

// Debugf writes debug level to output.
func (l Logger) Debugf(format string, args ...interface{}) {
	if l.Logger.WithFields {
		l.logger.WithFields(logrus.Fields(l.Fields)).Debugf(format, args...)
		return
	}

	l.logger.Debugf(format, args...)
}

// Infof writes info level to output.
func (l Logger) Infof(format string, args ...interface{}) {
	if l.Logger.WithFields {
		l.logger.WithFields(logrus.Fields(l.Fields)).Infof(format, args...)
		return
	}

	l.logger.Infof(format, args...)
}

// Warnf writes warn level to output.
func (l Logger) Warnf(format string, args ...interface{}) {
	if l.Logger.WithFields {
		l.logger.WithFields(logrus.Fields(l.Fields)).Warnf(format, args...)
		return
	}

	l.logger.Warnf(format, args...)
}

// Errorf writes error level to output.
func (l Logger) Errorf(format string, args ...interface{}) {
	if l.Logger.WithFields {
		l.logger.WithFields(logrus.Fields(l.Fields)).Errorf(format, args...)
		return
	}

	l.logger.Errorf(format, args...)
}

// WithFields supports fields.
func (l Logger) WithFields(fields models.LoggerFields) models.Loggerer {
	l.Logger.WithFields = true
	l.Logger.Fields = make(models.LoggerFields)

	for k, v := range fields {
		l.Logger.Fields[k] = v
	}

	return l
}
