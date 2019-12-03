/*
 * HomeWork-14: RabbitMQ sender
 * Created on 28.11.2019 22:45
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
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	eventsQueueName  = "events"
	eventIDField     = "event_id"
	eventOccursField = "occurs_at"
	eventAlertField  = "alert_every"
)

type publisher struct {
	conn    *amqp.Connection
	db      storage.DB
	ch      *amqp.Channel
	timeout time.Duration
	ctx     context.Context
	cancel  context.CancelFunc
	logger  loggers.Logger
}

func newPublisher(db storage.DB, dsn string, timeout time.Duration) (*publisher, error) {
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

	p := &publisher{
		conn:    conn,
		ch:      ch,
		db:      db,
		timeout: timeout,
		logger:  logger,
	}
	p.ctx, p.cancel = context.WithCancel(context.Background())

	return p, nil
}

func (p *publisher) close() error {
	errCh := p.ch.Close()
	errConn := p.conn.Close()
	if errCh != nil || errConn != nil {
		return fmt.Errorf("error close rabbit MQ: channel - %s, conn - %s", errCh, errConn)
	}
	p.logger.Info("Closed rabbit MQ")
	return nil
}

func (p *publisher) publish(event models.Event) error {
	q, err := p.ch.QueueDeclare(
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

	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	if err = p.ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		}); err != nil {
		return err
	}
	return nil
}

func (p *publisher) start() {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go p.worker()

	p.logger.Warn("Signal received: %s", <-shutdown)
	p.cancel()
}

func (p *publisher) worker() {
	alerts := make(map[uuid.UUID]time.Time)
OUTER:
	for {
		select {
		case <-p.ctx.Done():
			break OUTER
		case <-time.After(p.timeout):
			p.logger.Info("Publish worker tick: %s", p.timeout)
			events := p.db.GetAlertedEventsDB(p.ctx, time.Now())
			if len(events) == 0 {
				continue
			}

			for _, event := range events {
				if event.AlertEvery <= 0 {
					delete(alerts, event.ID)
					p.logger.WithFields(loggers.Fields{
						eventIDField:     event.ID.String(),
						eventOccursField: event.OccursAt,
						eventAlertField:  event.AlertEvery,
					}).Warn("Alerted skipped: every time <= 0")
					continue
				}

				if at, ok := alerts[event.ID]; ok {
					if at.Add(event.AlertEvery).After(time.Now()) {
						p.logger.WithFields(loggers.Fields{
							eventIDField:     event.ID.String(),
							eventOccursField: event.OccursAt,
							eventAlertField:  event.AlertEvery,
						}).Warn("Alerted skipped: delta every time < now()")
						continue
					}
				}

				if err := p.publish(event); err != nil {
					p.logger.WithFields(loggers.Fields{
						eventIDField: event.ID.String(),
					}).Error(err.Error())
					continue
				}
				alerts[event.ID] = time.Now()

				p.logger.WithFields(loggers.Fields{
					eventIDField:     event.ID.String(),
					eventOccursField: event.OccursAt,
					eventAlertField:  event.AlertEvery,
				}).Info("Alerted event published into queue")
				p.logger.Debug("Event body: %+v", event)
			}
		}
	}
	p.logger.Info("Publish worker ended")
}
