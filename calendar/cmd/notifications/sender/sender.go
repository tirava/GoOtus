/*
 * HomeWork-14: RabbitMQ receiver
 * Created on 30.11.2019 22:06
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/evakom/calendar/internal/domain/interfaces/storage"
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/evakom/calendar/internal/loggers"
	"github.com/streadway/amqp"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	eventsQueueName  = "events"
	eventIDField     = "event_id"
	eventOccursField = "occurs_at"
	userIDField      = "user_id"
	fromEmail        = "info@kirk.com"
	emailField       = "email"
)

type sender struct {
	conn   *amqp.Connection
	db     storage.DB
	ch     *amqp.Channel
	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
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

	p := &sender{
		conn:   conn,
		ch:     ch,
		db:     db,
		logger: logger,
	}
	p.ctx, p.cancel = context.WithCancel(context.Background())

	return p, nil
}

func (s *sender) close() error {
	errCh := s.ch.Close()
	errConn := s.conn.Close()
	if errCh != nil || errConn != nil {
		return fmt.Errorf("error close rabbit MQ: channel - %s, conn - %s", errCh, errConn)
	}
	s.logger.Info("Closed rabbit MQ")
	return nil
}

func (s *sender) start() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	if err := s.consume(); err != nil {
		log.Fatal(err)
	}

	s.logger.Warn("Signal received: %s", <-shutdown)
	s.cancel()
	s.wg.Wait()
}

func (s *sender) consume() error {
	q, err := s.ch.QueueDeclare(
		eventsQueueName, // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		return err
	}

	messages, err := s.ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	go func() {
	OUTER:
		for {
			select {
			case <-s.ctx.Done():
				break OUTER
			case msg := <-messages:
				s.logger.Info("Received message from queue")
				s.logger.Debug("Message body: %s", msg.Body)
				s.wg.Add(1)
				go s.parseAndSend(msg)
			}
		}
		s.logger.Info("Consume worker ended")
	}()

	return nil
}

func (s *sender) parseAndSend(msg amqp.Delivery) {
	defer s.wg.Done()

	event := models.Event{}
	if err := json.Unmarshal(msg.Body, &event); err != nil {
		s.logger.Error("Error parse message body: %s", msg.Body)
	}

	s.logger.WithFields(loggers.Fields{
		eventIDField:     event.ID.String(),
		eventOccursField: event.OccursAt,
		userIDField:      event.UserID,
	}).Info("Alerted event parsed from message")

	user, err := s.db.GetUserDB(s.ctx, event.UserID)
	if err != nil {
		s.logger.Error("Error get user data from DB: %s", err)
		return
	}

	if err := s.sendAlert(user, event); err != nil {
		s.logger.WithFields(loggers.Fields{
			eventIDField: event.ID.String(),
			userIDField:  event.UserID,
		}).Error("Error sending alert to user: %s", err)
	}
}

func (s *sender) sendAlert(user models.User, event models.Event) error {

	if err := sendEmail(
		fromEmail,
		event.Subject,
		event.Body,
		[]string{user.Email},
	); err != nil {
		return err
	}

	s.logger.WithFields(loggers.Fields{
		eventIDField: event.ID.String(),
		userIDField:  event.UserID,
	}).Info("Alerted event sent to user")

	event.AlertEvery = -1
	if err := s.db.EditEventDB(s.ctx, event); err != nil {
		return err
	}

	s.logger.WithFields(loggers.Fields{
		eventIDField: event.ID.String(),
		userIDField:  event.UserID,
		emailField:   user.Email,
	}).Info("Alerted event disabled for user")

	return nil
}
