/*
 * HomeWork-9: Integration tests
 * Created on 17.12.2019 17:24
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/evakom/calendar/internal/grpc/api"
	"github.com/evakom/calendar/tools"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"log"
	"time"
)

const dayLayout = "2006-01-02"

type eventsTest struct {
	req      *api.EventRequest
	resp     *api.EventResponse
	resps    *api.EventsResponse
	conn     *grpc.ClientConn
	client   api.CalendarServiceClient
	ctx      context.Context
	startDay *timestamp.Timestamp
}

func (t *eventsTest) start(interface{}) {
	var err error

	conf := tools.InitConfig("config.yml")
	t.conn, err = grpc.Dial(conf.ListenGRPC, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	t.client = api.NewCalendarServiceClient(t.conn)

	t.ctx = context.TODO()

	t.req = &api.EventRequest{
		Subject:    "GoDog list events",
		Body:       "HomeWork-9: Integration tests",
		Location:   "Moscow",
		Duration:   parseDuration("1h"),
		UserID:     "a7fdcee4-8a27-4200-8529-c5336c886f78",
		AlertEvery: parseDuration("-1ns"),
	}

	start := time.Now().Format(dayLayout)
	t.startDay = parseDateTime(start, dayLayout)
}

func (t *eventsTest) stop(interface{}, error) {
	if err := t.conn.Close(); err != nil {
		log.Println(err)
	}
}

func (t *eventsTest) iSendCreateEventToServiceAPIForCycleWithEventsForSameUserAndStepDaysForOccursAt(
	numEvents, stepDays int) error {

	for i := 0; i < numEvents; i++ {
		deltaDays := time.Duration(stepDays*i) * time.Hour * 24
		occursAt, err := ptypes.TimestampProto(time.Now().Add(deltaDays))
		if err != nil {
			return err
		}
		t.req.OccursAt = occursAt
		t.resp, err = t.client.CreateEvent(t.ctx, t.req)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *eventsTest) allAddedEventsWillBeReturnedByGetUserEventsForGivenUser(
	numEvents int) error {
	var err error

	t.resps, err = t.client.GetUserEvents(t.ctx, &api.ID{Id: t.req.GetUserID()})
	if err != nil {
		return err
	}

	actualEvents := len(t.resps.GetEvents())
	if actualEvents != numEvents {
		return fmt.Errorf("expected events: %d != actual events: %d",
			numEvents, actualEvents)
	}

	return nil
}

func (t *eventsTest) getErrorHasNoErrorsInTheseCases() error {
	if respsErr := t.resps.GetError(); respsErr != "" {
		return errors.New(respsErr)
	}

	return nil
}

func (t *eventsTest) iSendGetEventsForDayRequestWithCurrentDayToServiceAPI() error {
	var err error

	t.resps, err = t.client.GetEventsForDay(t.ctx, &api.Day{Day: t.startDay})
	if err != nil {
		return err
	}

	return nil
}

func (t *eventsTest) iGetEventsResponseWithEventInItWithOccursAtInCurrentDay(numEvents int) error {
	actualEvents := len(t.resps.GetEvents())
	if actualEvents != numEvents {
		return fmt.Errorf("expected events: %d != actual events: %d",
			numEvents, actualEvents)
	}

	return nil
}

func (t *eventsTest) getErrorReturnsNoErrors() error {
	return t.getErrorHasNoErrorsInTheseCases()
}

func (t *eventsTest) iSendGetEventsForWeekRequestWithCurrentDayToServiceAPI() error {
	var err error

	t.resps, err = t.client.GetEventsForWeek(t.ctx, &api.Day{Day: t.startDay})
	if err != nil {
		return err
	}

	return nil
}

func (t *eventsTest) iGetEventsResponseWithEventsInItWithOccursAtInNearWeek(numEvents int) error {
	return t.iGetEventsResponseWithEventInItWithOccursAtInCurrentDay(numEvents)
}

func (t *eventsTest) iSendGetEventsForMonthRequestWithCurrentDayToServiceAPI() error {
	var err error

	t.resps, err = t.client.GetEventsForMonth(t.ctx, &api.Day{Day: t.startDay})
	if err != nil {
		return err
	}

	return nil
}

func (t *eventsTest) iGetEventsResponseWithEventsInItWithOccursAtInNearMonth(numEvents int) error {
	return t.iGetEventsResponseWithEventInItWithOccursAtInCurrentDay(numEvents)
}

// FeatureContextListEvents implements test suite.
func FeatureContextListEvents(s *godog.Suite) {
	tests := new(eventsTest)
	s.BeforeScenario(tests.start)

	s.Step(`^I send CreateEvent to service API for cycle with (\d+) events for same user and step (\d+) days for OccursAt$`,
		tests.iSendCreateEventToServiceAPIForCycleWithEventsForSameUserAndStepDaysForOccursAt)
	s.Step(`^all (\d+) added events will be returned by GetUserEvents for given user$`,
		tests.allAddedEventsWillBeReturnedByGetUserEventsForGivenUser)
	s.Step(`^GetError has no errors in these cases$`,
		tests.getErrorHasNoErrorsInTheseCases)
	s.Step(`^I send GetEventsForDay request with current day to service API$`,
		tests.iSendGetEventsForDayRequestWithCurrentDayToServiceAPI)
	s.Step(`^I get EventsResponse with (\d+) event in it with OccursAt in current day$`,
		tests.iGetEventsResponseWithEventInItWithOccursAtInCurrentDay)
	s.Step(`^GetError returns no errors$`,
		tests.getErrorReturnsNoErrors)
	s.Step(`^I send GetEventsForWeek request with current day to service API$`,
		tests.iSendGetEventsForWeekRequestWithCurrentDayToServiceAPI)
	s.Step(`^I get EventsResponse with (\d+) events in it with OccursAt in near week$`,
		tests.iGetEventsResponseWithEventsInItWithOccursAtInNearWeek)
	s.Step(`^I send GetEventsForMonth request with current day to service API$`,
		tests.iSendGetEventsForMonthRequestWithCurrentDayToServiceAPI)
	s.Step(`^I get EventsResponse with (\d+) events in it with OccursAt in near month$`,
		tests.iGetEventsResponseWithEventsInItWithOccursAtInNearMonth)

	s.AfterScenario(tests.stop)
}
