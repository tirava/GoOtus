/*
 * HomeWork-9: Calendar protobuf preparation
 * Created on 30.10.2019 15:18
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package models implements base models.
package models

import (
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
)

// Fields is the log fields type.
type Fields map[string]interface{}

// Logger is the base struct for all loggers.
type Logger struct {
	logger     *log.Logger
	Fields     Fields
	withFields bool
}

var lg Logger

// GetLogger returns global logger.
func (l Logger) GetLogger() Logger {
	return lg
}

// NewLogger inits logger instance.
func NewLogger(level string, output io.Writer) {
	lg = Logger{logger: log.New()}
	if level == "none" {
		lg.logger.SetOutput(ioutil.Discard)
	} else {
		lg.logger.SetOutput(output)
	}
	switch level {
	case "debug":
		lg.logger.SetLevel(log.DebugLevel)
	case "info":
		lg.logger.SetLevel(log.InfoLevel)
	case "warn":
		lg.logger.SetLevel(log.WarnLevel)
	default:
		lg.logger.SetLevel(log.ErrorLevel)
	}
}

// Debug writes debug level to output.
func (l Logger) Debug(format string, args ...interface{}) {
	if l.withFields {
		l.logger.WithFields(log.Fields(l.Fields)).Debugf(format, args...)
		return
	}
	l.logger.Debugf(format, args...)
}

// Info writes info level to output.
func (l Logger) Info(format string, args ...interface{}) {
	if l.withFields {
		l.logger.WithFields(log.Fields(l.Fields)).Infof(format, args...)
		return
	}
	l.logger.Infof(format, args...)
}

// Warn writes warn level to output.
func (l Logger) Warn(format string, args ...interface{}) {
	if l.withFields {
		l.logger.WithFields(log.Fields(l.Fields)).Warnf(format, args...)
		return
	}
	l.logger.Warnf(format, args...)
}

// Error writes error level to output.
func (l Logger) Error(format string, args ...interface{}) {
	if l.withFields {
		l.logger.WithFields(log.Fields(l.Fields)).Errorf(format, args...)
		return
	}
	l.logger.Errorf(format, args...)
}

// WithFields support fields.
func (l Logger) WithFields() Logger {
	l.withFields = true
	return l
}
