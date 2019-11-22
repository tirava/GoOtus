/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 27.10.2019 12:32
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"context"
	"flag"
	"github.com/evakom/calendar/internal/domain/calendar"
	"github.com/evakom/calendar/tools"
	"github.com/evakom/calendar/website"
)

func main() {
	configFile := flag.String("config", "config.yml", "path to config file")
	flag.Parse()

	conf := tools.InitConfig(*configFile)

	logFile := tools.InitLogger(conf)
	defer logFile.Close()

	db := tools.InitDB(context.TODO(), conf.DBType, conf.DSN)

	cal := calendar.NewCalendar(db)

	website.StartWebsite(conf.ListenHTTP, cal)
}
