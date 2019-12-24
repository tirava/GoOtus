/*
 * HomeWork-14: RabbitMQ receiver
 * Created on 28.11.2019 22:20
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"context"
	"flag"
	"log"

	"github.com/evakom/calendar/tools"
)

func main() {
	configFile := flag.String("config", "config.yml", "path to config file")
	flag.Parse()

	conf := tools.InitConfig(*configFile)

	logFile := tools.InitLogger(conf)
	defer logFile.Close()

	db := tools.InitDB(context.TODO(), conf.DBType, conf.DSN)

	sender, err := newSender(db, conf.RabbitMQ)
	if err != nil {
		log.Fatal(err)
	}

	promet := newPrometheus(":9102")
	sender.promet = promet
	promet.start()

	sender.start()

	if err := sender.close(); err != nil {
		log.Println("Error close RabbitMQ:", err)
	}
	if err := db.CloseDB(); err != nil {
		log.Println("Error close DB:", err)
	}

	promet.stop()
}
