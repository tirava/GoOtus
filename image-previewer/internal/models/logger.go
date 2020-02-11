package models

// Loggerer is the main interface for loggers.
type Loggerer interface {
	GetLogger() Loggerer
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	WithFields(LoggerFields) Loggerer
}

// LoggerFields is the log fields type.
type LoggerFields map[string]interface{}

// Logger is the base struct for all loggers.
type Logger struct {
	Fields     LoggerFields
	WithFields bool
}
