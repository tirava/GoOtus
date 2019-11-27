/*
 * HomeWork-12: gRPC server
 * Created on 24.11.2019 12:45
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"context"
	"flag"
	"github.com/evakom/calendar/internal/domain/calendar"
	"github.com/evakom/calendar/internal/grpc/api"
	"github.com/evakom/calendar/tools"
	"log"
)

func main() {

	configFile := flag.String("config", "config.yml", "path to config file")
	flag.Parse()

	conf := tools.InitConfig(*configFile)

	logFile := tools.InitLogger(conf)
	defer logFile.Close()

	db := tools.InitDB(context.TODO(), conf.DBType, conf.DSN)
	cal := calendar.NewCalendar(db)

	cs := api.NewCalendarServer(cal)
	cs.StartGRPCServer(conf.ListenGRPC)

	if err := db.CloseDB(); err != nil {
		log.Println("Error close DB:", err)
	}
}
