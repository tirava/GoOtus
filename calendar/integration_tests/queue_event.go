/*
 * HomeWork-9: Integration tests
 * Created on 18.12.2019 13:27
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/evakom/calendar/internal/grpc/api"
	"github.com/evakom/calendar/tools"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"
)

const (
	eventsQueueName = "events"
	dayFormat       = "2006-01-02"
)

type alertTest struct {
	connAmpq *amqp.Connection
	ch       *amqp.Channel
	sync.RWMutex
	messages    [][]byte
	req         *api.EventRequest
	resp        *api.EventResponse
	connGrpc    *grpc.ClientConn
	client      api.CalendarServiceClient
	ctx         context.Context
	waitSched   time.Duration
	wasQueueErr bool
}

func (t *alertTest) startConsume(interface{}) {
	var err error
	t.messages = make([][]byte, 0)
	//t.stop = make(chan bool)

	conf := tools.InitConfig("config.yml")

	t.connAmpq, err = amqp.Dial(conf.RabbitMQ)
	if err != nil {
		log.Fatal(err)
	}

	t.ch, err = t.connAmpq.Channel()
	if err != nil {
		log.Fatal(err)
	}

	t.waitSched, err = time.ParseDuration(conf.PubTimeout)
	if err != nil {
		log.Fatal(err)
	}

	t.connGrpc, err = grpc.Dial(conf.ListenGRPC, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	t.client = api.NewCalendarServiceClient(t.connGrpc)

	t.ctx = context.TODO()

	start := time.Now().Format(dayFormat)

	t.req = &api.EventRequest{
		OccursAt:   parseDateTime(start, dayFormat),
		Subject:    "GoDog alert event",
		Body:       "HomeWork-9: Integration tests",
		Location:   "Moscow",
		Duration:   parseDuration("1h"),
		UserID:     "a7fdcee4-8a27-4200-8529-c5336c886f79",
		AlertEvery: parseDuration("1m"),
	}
}

func (t *alertTest) stopConsume(interface{}, error) {
	errCh := t.ch.Close()
	errConn := t.connAmpq.Close()

	if errCh != nil || errConn != nil {
		log.Println(errCh, errConn)
	}

	t.messages = nil
}

func (t *alertTest) iCreateEventWithEventRequestToServiceAPIWithOccursAtNow() error {
	var err error

	t.resp, err = t.client.CreateEvent(t.ctx, t.req)
	if err != nil {
		return err
	}

	return nil
}

func (t *alertTest) addedEventWillBeScheduledIntoMessageQueue() error {
	// it will be done by scheduler
	return nil
}

func (t *alertTest) getErrorHasNoError() error {
	if respErr := t.resp.GetError(); respErr != "" {
		return errors.New(respErr)
	}

	return nil
}

func (t *alertTest) iConsumeMessageQueue() error {

	q, err := t.ch.QueueDeclare(eventsQueueName, false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	events, err := t.ch.Consume(q.Name, "", true, true, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for event := range events {
			t.Lock()
			t.messages = append(t.messages, event.Body)
			t.Unlock()
		}
	}()

	fmt.Printf("Wait %s for scheduler send event...", t.waitSched.String())
	time.Sleep(t.waitSched)

	t.RLock()
	defer t.RUnlock()
	if len(t.messages) != 1 {
		return fmt.Errorf("event in queue must be one, but received: %d",
			len(t.messages))
	}

	return nil
}

func (t *alertTest) iGetEventWithCorrectTestUserID() error {
	t.RLock()
	defer t.RUnlock()

	event := models.Event{}
	if err := json.Unmarshal(t.messages[0], &event); err != nil {
		return err
	}

	if event.UserID.String() != t.req.UserID {
		t.wasQueueErr = true
		return fmt.Errorf("user id from database: %s != user id from queue: %s",
			t.resp.GetEvent().UserID, event.UserID.String())
	}

	return nil
}

func (t *alertTest) willBeReadyToSendMessageRotThisUser() error {
	if t.wasQueueErr {
		return errors.New("alert email not sent to user")
	}

	fmt.Printf("Sent fake email to User id: %s\n", t.resp.GetEvent().UserID)

	return nil
}

// FeatureContextQueueEvent implements test suite.
func FeatureContextQueueEvent(s *godog.Suite) {
	test := new(alertTest)
	s.BeforeScenario(test.startConsume)

	s.Step(`^I CreateEvent with EventRequest to service API with OccursAt = Now$`,
		test.iCreateEventWithEventRequestToServiceAPIWithOccursAtNow)
	s.Step(`^added event will be scheduled into message queue$`,
		test.addedEventWillBeScheduledIntoMessageQueue)
	s.Step(`^GetError has no error$`,
		test.getErrorHasNoError)
	s.Step(`^I consume message queue$`,
		test.iConsumeMessageQueue)
	s.Step(`^I get event with correct test user id$`,
		test.iGetEventWithCorrectTestUserID)
	s.Step(`^will be ready to send message rot this user$`,
		test.willBeReadyToSendMessageRotThisUser)

	s.AfterScenario(test.stopConsume)
}
