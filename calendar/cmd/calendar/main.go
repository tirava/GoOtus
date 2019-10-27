/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 27.10.2019 12:32
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"github.com/evakom/calendar/pkg/calendar"
	"log"
	"os"
)

// Constants
const (
	EnvCalendarConfigPath  = "CALENDAR_CONFIG_PATH"
	FileCalendarConfigPath = "./configs/calendar.yml"
)

func main() {

	confPath := os.Getenv(EnvCalendarConfigPath)

	if confPath == "" {
		confPath = FileCalendarConfigPath
	}

	conf := calendar.NewConfig(confPath)
	if err := conf.ReadParameters(); err != nil {
		log.Fatalln(err)
	}

	db := calendar.NewDB(conf.DBType)
	if db == nil {
		log.Fatalf("unsupported DB type: %s", conf.DBType)
	}

	if db.MapDB != nil {
		calendar.PrintTestData(db.MapDB)
	} //else if db.PostgresDB != nil {
	//calendar.PrintTestData(db.PostgresDB)
	//}
}
