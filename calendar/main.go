/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 27.10.2019 12:32
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"github.com/evakom/calendar/internal/domain/interfaces"
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/evakom/calendar/internal/pkg/calendar"
	"log"
	"os"
)

// Constants
const (
	EnvCalendarConfigPath  = "CALENDAR_CONFIG_PATH"
	FileCalendarConfigPath = "./internal/configs/calendar.yml"
)

func main() {

	confPath := os.Getenv(EnvCalendarConfigPath)

	if confPath == "" {
		confPath = FileCalendarConfigPath
	}

	conf := models.NewConfig(confPath)
	if err := conf.ReadParameters(); err != nil {
		log.Fatalln(err)
	}

	db := interfaces.NewDB(conf.DBType)
	if db == nil {
		log.Fatalf("unsupported DB type: %s\n", conf.DBType)
	}

	calendar.PrintTestData(db)
}
