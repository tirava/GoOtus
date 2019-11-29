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
	"os"
	"os/signal"
	"syscall"
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

	//err = sender.publish("111", "222")
	//if err != nil {
	//	log.Fatal(err)
	//}
	go func() {

	}()

	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	sender.logger.Warn("Signal received: %s", <-shutdown)

	if err := sender.close(); err != nil {
		log.Println("Error close RabbitMQ:", err)
	}
	if err := db.CloseDB(); err != nil {
		log.Println("Error close DB:", err)
	}
}
