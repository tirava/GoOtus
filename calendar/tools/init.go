/*
 * HomeWork-9: Calendar protobuf preparation
 * Created on 03.11.2019 13:27
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package tools implements inits and tools of various subsystems.
package tools

import (
	"github.com/evakom/calendar/internal/configs"
	"github.com/evakom/calendar/internal/dbs"
	"github.com/evakom/calendar/internal/domain/interfaces"
	"github.com/evakom/calendar/internal/domain/models"
	"log"
	"os"
)

// Constants
const (
	envCalendarConfigPath = "CALENDAR_CONFIG_PATH"
)

// InitConfig inits config params.
func InitConfig(configFile string) configs.Config {
	confPath := os.Getenv(envCalendarConfigPath)
	if confPath == "" {
		confPath = configFile
	}
	conf, err := configs.NewConfig(confPath)
	if err != nil {
		log.Fatalln(err)
	}
	return conf
}

// InitLogger return logger file.
func InitLogger(conf configs.Config) *os.File {
	logFile, err := os.OpenFile(conf.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error open log file '%s', error: %s", conf.LogFile, err)
	}
	models.NewLogger(conf.LogLevel, logFile)
	return logFile
}

// InitDB inits db interface.
func InitDB(dbType string) interfaces.DB {
	db, err := dbs.NewDB(dbType)
	if db == nil {
		log.Fatalf("unsupported DB type: %s\n", dbType)
	}
	if err != nil {
		log.Fatalf("Open DB: %s, error: %s \n", dbType, err)
	}
	return db
}
