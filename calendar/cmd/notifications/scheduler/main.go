/*
 * HomeWork-14: RabbitMQ sender
 * Created on 28.11.2019 22:20
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"context"
	"flag"
	"github.com/evakom/calendar/tools"
	"log"
	"time"
)

func main() {
	configFile := flag.String("config", "config.yml", "path to config file")
	flag.Parse()

	conf := tools.InitConfig(*configFile)

	timeout, err := time.ParseDuration(conf.PubTimeout)
	if err != nil {
		log.Fatal(err)
	}

	if timeout < 100*time.Millisecond {
		timeout = 100 * time.Millisecond
	}

	logFile := tools.InitLogger(conf)
	defer logFile.Close()

	db := tools.InitDB(context.TODO(), conf.DBType, conf.DSN)

	publisher, err := newScheduler(db, conf.RabbitMQ, timeout)
	if err != nil {
		log.Fatal(err)
	}

	publisher.start()

	if err := publisher.close(); err != nil {
		log.Println("Error close RabbitMQ:", err)
	}
	if err := db.CloseDB(); err != nil {
		log.Println("Error close DB:", err)
	}
}
