/*
 * HomeWork-14: RabbitMQ sender
 * Created on 28.11.2019 22:45
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"fmt"
	"github.com/evakom/calendar/internal/domain/interfaces/storage"
	"github.com/evakom/calendar/internal/loggers"
	"github.com/streadway/amqp"
)

type sender struct {
	conn   *amqp.Connection
	db     storage.DB
	ch     *amqp.Channel
	logger loggers.Logger
}

func newSender(db storage.DB, dsn string) (*sender, error) {
	conn, err := amqp.Dial(dsn)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	logger := loggers.GetLogger()
	logger.Info("Connected to rabbit MQ")
	return &sender{
		conn:   conn,
		ch:     ch,
		db:     db,
		logger: logger,
	}, nil
}

func (s *sender) close() error {
	errCh := s.ch.Close()
	errConn := s.conn.Close()
	if errCh != nil || errConn != nil {
		return fmt.Errorf("error close rabbit MQ: %s, %s", errCh, errConn)
	}
	s.logger.Info("Closed rabbit MQ")
	return nil
}
