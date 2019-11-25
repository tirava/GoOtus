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
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"os"
	"path/filepath"
	"time"
)

const tsLayout = "2006-01-02 15:04:05"

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
)

func init() {

	fileName := filepath.Base(os.Args[0])
	flag.Usage = func() {
		fmt.Printf("Call server on custom host:port: %s -server host:port -method ...\n", fileName)
		fmt.Printf("Create event:       %s -method create_event -user_id uuid "+
			"[-occurs_at 'date time'] [-duration duration] "+
			"[-subject 'subject'] [-body 'body'] [-location 'location']\n", fileName)
		fmt.Printf("Get event:          %s -method get_event -event_id uuid\n", fileName)
		fmt.Printf("Get user events:    %s -method get_user_events -user_id uuid\n", fileName)
		flag.PrintDefaults()
	}

	flag.StringVar(&server, "server", "localhost:50051", "server host:port")
	flag.StringVar(&method, "method", "get_event", "call method")
	flag.StringVar(&uid, "user_id", uuid.Nil.String(), "owner uuid")
	flag.StringVar(&eid, "event_id", uuid.Nil.String(), "event uuid")
	flag.StringVar(&occursAt, "occurs_at", time.Now().Format(tsLayout),
		"date and time when event occurs")
	flag.StringVar(&duras, "duration", "24h", "event duration (sec, min, hours)")
	flag.StringVar(&subject, "subject", "", "event subject (title)")
	flag.StringVar(&body, "body", "", "event body (description)")
	flag.StringVar(&location, "location", "", "event location (where)")
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := api.NewCalendarServiceClient(conn)

	occurs, err := parseDateTime(occursAt)
	if err != nil {
		log.Fatal(err)
	}
	durat, err := parseDuration(duras)
	if err != nil {
		log.Fatal(err)
	}

	req := &api.EventRequest{
		OccursAt: occurs,
		Subject:  subject,
		Body:     body,
		Location: location,
		Duration: durat,
		UserID:   uid,
	}
	eId := &api.ID{Id: eid}
	uId := &api.ID{Id: uid}
	resp := &api.EventResponse{}
	resps := &api.EventsResponse{}
	ctx := context.TODO()
	needUsage := false

	switch method {
	case "create_event":
		if uid == "" || uid == uuid.Nil.String() {
			needUsage = true
			break
		}
		resp, err = client.CreateEvent(ctx, req)
	case "get_event":
		if eid == "" || eid == uuid.Nil.String() {
			needUsage = true
			break
		}
		resp, err = client.GetEvent(ctx, eId)
	case "get_user_events":
		if uid == "" || uid == uuid.Nil.String() {
			needUsage = true
			break
		}
		resps, err = client.GetUserEvents(ctx, uId)
	default:
		needUsage = true
	}

	if needUsage {
		flag.Usage()
		os.Exit(2)
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

func parseDateTime(s string) (*timestamp.Timestamp, error) {
	t, err := time.Parse(tsLayout, s)
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
