/*
 * HomeWork-12: gRPC client
 * Created on 24.11.2019 13:44
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/evakom/calendar/internal/grpc/api"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	tsLayout  = "2006-01-02 15:04:05"
	dayLayout = "2006-01-02"
)

var (
	server   string
	method   string
	uid      string
	eid      string
	occursAt string
	duras    string
	subject  string
	body     string
	location string
	startDay string
	alert    string
)

func init() {

	fileName := filepath.Base(os.Args[0])
	flag.Usage = func() {
		fmt.Printf("Call server on custom host:port: %s -server host:port -method ...\n", fileName)
		fmt.Printf("Create event:         %s -method create_event -user_id uuid "+
			"[-occurs_at 'date time'] [-duration duration] "+
			"[-subject 'subject'] [-body 'body'] [-location 'location']\n", fileName)
		fmt.Printf("Update event:         %s -method update_event -event_id uuid "+
			"[-occurs_at 'date time'] [-duration duration] [-alert_every duration]"+
			"[-subject 'subject'] [-body 'body'] [-location 'location']\n", fileName)
		fmt.Printf("Get event:            %s -method get_event -event_id uuid\n", fileName)
		fmt.Printf("Delete event:         %s -method del_event -event_id uuid\n", fileName)
		fmt.Printf("Get user events:      %s -method get_user_events -user_id uuid\n", fileName)
		fmt.Printf("Get events for day:   %s -method get_events_day -start_day date\n", fileName)
		fmt.Printf("Get events for week:  %s -method get_events_week -start_day date\n", fileName)
		fmt.Printf("Get events for month: %s -method get_events_month -start_day date\n", fileName)
		flag.PrintDefaults()
	}

	flag.StringVar(&server, "server", "localhost:50051", "server host:port")
	flag.StringVar(&method, "method", "get_event", "call method")
	flag.StringVar(&uid, "user_id", "", "owner uuid")
	flag.StringVar(&eid, "event_id", "", "event uuid")
	flag.StringVar(&occursAt, "occurs_at", time.Time{}.Format(tsLayout),
		"date and time when event occurs")
	flag.StringVar(&duras, "duration", "1h", "event duration (sec, min, hours)")
	flag.StringVar(&subject, "subject", "", "event subject (title)")
	flag.StringVar(&body, "body", "", "event body (description)")
	flag.StringVar(&location, "location", "", "event location (where)")
	flag.StringVar(&startDay, "start_day", time.Now().Format(dayLayout),
		"start date when events will occur")
	flag.StringVar(&alert, "alert_every", "15m", "event alerts every (sec, min, hours)")
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := api.NewCalendarServiceClient(conn)

	occurs, err := parseDateTime(occursAt, tsLayout)
	if err != nil {
		log.Fatal(err)
	}
	durat, err := parseDuration(duras)
	if err != nil {
		log.Fatal(err)
	}
	every, err := parseDuration(alert)
	if err != nil {
		log.Fatal(err)
	}
	start, err := parseDateTime(startDay, dayLayout)
	if err != nil {
		log.Fatal(err)
	}

	req := &api.EventRequest{
		ID:         eid,
		OccursAt:   occurs,
		Subject:    subject,
		Body:       body,
		Location:   location,
		Duration:   durat,
		UserID:     uid,
		AlertEvery: every,
	}

	eID := &api.ID{Id: eid}
	uID := &api.ID{Id: uid}
	resp := &api.EventResponse{}
	resps := &api.EventsResponse{}
	day := &api.Day{Day: start}
	ctx := context.TODO()

	switch method {
	case "create_event":
		resp, err = client.CreateEvent(ctx, req)
	case "update_event":
		resp, err = client.UpdateEvent(ctx, req)
	case "get_event":
		resp, err = client.GetEvent(ctx, eID)
	case "del_event":
		resp, err = client.DeleteEvent(ctx, eID)
	case "get_user_events":
		resps, err = client.GetUserEvents(ctx, uID)
	case "get_events_day":
		resps, err = client.GetEventsForDay(ctx, day)
	case "get_events_week":
		resps, err = client.GetEventsForWeek(ctx, day)
	case "get_events_month":
		resps, err = client.GetEventsForMonth(ctx, day)
	}

	if err != nil {
		log.Fatal(err)
	}

	if resp.GetError() != "" {
		log.Fatal(resp.GetError())
	}
	log.Println("Event:", resp.GetEvent())

	if resps.GetError() != "" {
		log.Fatal(resps.GetError())
	}
	log.Println("Events:", resps.GetEvents())

	if err := conn.Close(); err != nil {
		log.Fatal(err)
	}
}

func parseDateTime(s, layout string) (*timestamp.Timestamp, error) {
	t, err := time.Parse(layout, s)
	if err != nil {
		return nil, err
	}
	ts, err := ptypes.TimestampProto(t)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

func parseDuration(s string) (*duration.Duration, error) {
	d, err := time.ParseDuration(s)
	if err != nil {
		return nil, err
	}
	dr := ptypes.DurationProto(d)
	return dr, nil
}
