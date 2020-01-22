/*
 * Project: Image Previewer
 * Created on 20.01.2020 23:01
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package logstd implements std log logger.
package logstd

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"gitlab.com/tirava/image-previewer/internal/models"
)

// Logger is the base struct for std logger.
type Logger struct {
	logger *log.Logger
	models.Logger
	prefix map[string]string
	level  string
	fields interface{}
}

// NewLogger returns new logger.
func NewLogger(level string, output io.Writer) (Logger, error) {
	if level == "none" {
		output = ioutil.Discard
	}

	lg := Logger{
		logger: log.New(output, "", log.Ldate|log.Ltime),
		prefix: make(map[string]string, 4),
		level:  level,
	}

	lg.prefix["debug"] = "[DEBU] "
	lg.prefix["info"] = "[INFO] "
	lg.prefix["warn"] = "[WARN] "
	lg.prefix["error"] = "[ERRO] "

	return lg, nil
}

// GetLogger returns global logger.
func (l Logger) GetLogger() models.Loggerer {
	return l
}

func (l Logger) xPrintf(format string, args ...interface{}) {
	if l.Logger.WithFields {
		args = append(args, l.fields)
	}

	l.logger.Printf(format, args...)
}

// Debugf writes debug level to output.
func (l Logger) Debugf(format string, args ...interface{}) {
	l.logger.SetPrefix(l.prefix["debug"])

	l.xPrintf(format, args...)
}

// Infof writes info level to output.
func (l Logger) Infof(format string, args ...interface{}) {
	l.logger.SetPrefix(l.prefix["info"])

	l.xPrintf(format, args...)
}

// Warnf writes warn level to output.
func (l Logger) Warnf(format string, args ...interface{}) {
	l.logger.SetPrefix(l.prefix["warn"])

	l.xPrintf(format, args...)
}

// Errorf writes error level to output.
func (l Logger) Errorf(format string, args ...interface{}) {
	l.logger.SetPrefix(l.prefix["error"])

	l.xPrintf(format, args...)
}

// WithFields supports fields.
func (l Logger) WithFields(fields models.LoggerFields) models.Loggerer {
	l.Logger.WithFields = true
	sb := strings.Builder{}

	for k, v := range fields {
		sb.WriteString(fmt.Sprintf(" %s=%s", k, v))
	}

	l.fields = sb.String()

	return l
}
