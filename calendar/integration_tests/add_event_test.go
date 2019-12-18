/*
 * HomeWork-9: Integration tests
 * Created on 16.12.2019 10:44
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/DATA-DOG/godog"
	bizErr "github.com/evakom/calendar/internal/domain/errors"
	"github.com/evakom/calendar/internal/grpc/api"
	"github.com/evakom/calendar/tools"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
)

const tsLayout = "2006-01-02 15:04:05"

type eventTest struct {
	req    *api.EventRequest
	resp   *api.EventResponse
	conn   *grpc.ClientConn
	client api.CalendarServiceClient
	ctx    context.Context
	lastID string
	failID string
}

func (t *eventTest) start(interface{}) {
	var err error
	conf := tools.InitConfig("config.yml")
	t.conn, err = grpc.Dial(conf.ListenGRPC, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	t.client = api.NewCalendarServiceClient(t.conn)

	t.ctx = context.TODO()

	t.req = &api.EventRequest{
		OccursAt:   parseDateTime("2019-12-16 12:36:55", tsLayout),
		Subject:    "GoDog added event",
		Body:       "HomeWork-9: Integration tests",
		Location:   "Moscow",
		Duration:   parseDuration("1h"),
		UserID:     "a7fdcee4-8a27-4200-8529-c5336c886f77",
		AlertEvery: parseDuration("-1ns"),
	}
}

func (t *eventTest) stop(interface{}, error) {
	if err := t.conn.Close(); err != nil {
		log.Println(err)
	}
}

func (t *eventTest) iSendCreateEventWithEventRequestToServiceAPI() error {
	var err error

	t.resp, err = t.client.CreateEvent(t.ctx, t.req)
	if err != nil {
		return err
	}

	return nil
}

func (t *eventTest) addedEventWillBeReturnedByGetEventWithIDOfTheEvent() error {
	id := t.resp.GetEvent().Id
	_, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return nil
}

func (t *eventTest) getErrorHasNoErrorsInBothCases() error {
	if respErr := t.resp.GetError(); respErr != "" {
		return errors.New(respErr)
	}

	return nil
}

func (t *eventTest) iSendGetEventRequestWithEventIDToServiceAPI() error {
	var err error
	t.lastID = t.resp.GetEvent().Id

	t.resp, err = t.client.GetEvent(t.ctx, &api.ID{Id: t.lastID})
	if err != nil {
		return err
	}

	return nil
}

func (t *eventTest) iGetEventResponseWithIDOfTheEvent() error {
	respID := t.resp.GetEvent().Id
	if respID != t.lastID {
		return fmt.Errorf("request eventID: %s != response eventID: %s",
			t.lastID, respID)
	}

	return nil
}

func (t *eventTest) getErrorHasNoErrors() error {
	return t.getErrorHasNoErrorsInBothCases()
}

func (t *eventTest) iSendGetEventRequestWithNonExistingEventIDToServiceAPI() error {
	var err error
	t.failID = "4bfe7be9-6b28-4f9e-8545-a315217227f5"

	t.resp, err = t.client.GetEvent(t.ctx, &api.ID{Id: t.failID})
	if err != nil {
		return err
	}

	return nil
}

func (t *eventTest) iGetEventResponseWithErrorCodeEventNotFound() error {
	respErr := t.resp.GetError()
	if respErr != bizErr.ErrEventNotFound.Error() {
		return fmt.Errorf("expected error: %s but got error: %s",
			bizErr.ErrEventNotFound, respErr)
	}

	return nil
}

// FeatureContextAddEvent implements test suite.
func FeatureContextAddEvent(s *godog.Suite) {
	test := new(eventTest)
	s.BeforeScenario(test.start)

	s.Step(`^I send CreateEvent with EventRequest to service API$`,
		test.iSendCreateEventWithEventRequestToServiceAPI)
	s.Step(`^added event will be returned by GetEvent with id of the event$`,
		test.addedEventWillBeReturnedByGetEventWithIDOfTheEvent)
	s.Step(`^GetError has no errors in both cases$`,
		test.getErrorHasNoErrorsInBothCases)
	s.Step(`^I send GetEvent request with event id to service API$`,
		test.iSendGetEventRequestWithEventIDToServiceAPI)
	s.Step(`^I get EventResponse with id of the event$`,
		test.iGetEventResponseWithIDOfTheEvent)
	s.Step(`^GetError has no errors$`,
		test.getErrorHasNoErrors)
	s.Step(`^I send GetEvent request with non existing event id to service API$`,
		test.iSendGetEventRequestWithNonExistingEventIDToServiceAPI)
	s.Step(`^I get EventResponse with error code \'Event not found\'$`,
		test.iGetEventResponseWithErrorCodeEventNotFound)

	s.AfterScenario(test.stop)
}
