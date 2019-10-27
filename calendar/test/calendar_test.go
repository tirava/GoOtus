/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 24.10.2019 21:15
 * Copyright (c) 2019 - Eugene Klimov
 */

package test

import (
	"github.com/evakom/calendar/pkg/calendar"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"reflect"
	"strings"
	"testing"
)

func TestNewEvent(t *testing.T) {
	e1 := calendar.NewEvent().GetId()
	e2 := calendar.NewEvent().GetId()
	if e2 == e1 {
		t.Errorf("'id1 = %v' same as 'id2 = %v'", e1, e2)
	}
}

func TestProto(t *testing.T) {
	e1 := &calendar.Event{}
	e1.Subject = "testik"
	out, err := proto.Marshal(e1)
	if err != nil {
		t.Fatalf("Failed to encode event: %v, error: %v", e1, err)
	}
	e2 := &calendar.Event{}
	err = proto.Unmarshal(out, e2)
	if err != nil {
		t.Fatalf("Failed to parse event: %v, error: %v", out, err)
	}
	e2.XXX_sizecache = 8 // Subject length + 1
	if !reflect.DeepEqual(e1, e2) {
		t.Errorf("Event1 not equal Event2 after Marshal/Unmarshal:\n%#v\n%#v", e1, e2)
	}
}

func TestAddEvent(t *testing.T) {
	events := createNewDB()
	e := *calendar.NewEvent()
	e.Subject = "222222222222222222222"
	e.Body = "3333333333333333333"
	_ = events.AddEvent(e)
	e = *calendar.NewEvent()
	e.Duration = 555
	_ = events.AddEvent(e)
	l := len(events.Events)
	if l != 2 {
		t.Errorf("After adding 2 events to MapDB length != 2, actual length = %d", l)
	}
}

func TestGetEvent(t *testing.T) {
	events := createNewDB()
	e1 := *calendar.NewEvent()
	_ = events.AddEvent(e1)
	e2 := *calendar.NewEvent()
	_ = events.AddEvent(e2)
	e3, _ := events.GetEvent(e1.Id)
	if !reflect.DeepEqual(e1, e3) {
		t.Errorf("Event1 not equal Event3 after get from DB:\n%#v\n%#v", e1, e3)
	}
}

func TestEditEvent(t *testing.T) {
	events := createNewDB()
	e1 := *calendar.NewEvent()
	_ = events.AddEvent(e1)

	e2, _ := events.GetEvent(e1.Id)
	e2.Subject = "11111111111111111"
	e2.Body = "22222222222222222"
	_ = events.EditEvent(e2)

	e3, _ := events.GetEvent(e1.Id)
	if e3.Subject != e2.Subject || e3.Body != e2.Body {
		t.Errorf("Event1 not properly updated in the DB:\nExpected: %+v\nActual: %+v", e2, e3)
	}

	t2 := ptypes.TimestampString(e2.UpdatedAt)
	t3 := ptypes.TimestampString(e3.UpdatedAt)
	if strings.Compare(t2, t3) >= 0 {
		t.Errorf("Event1 updated time not correct in the DB:\nOld time: %v\nNew time: %v", t2, t3)
	}
}

func TestDelEvent(t *testing.T) {
	events := createNewDB()
	e := *calendar.NewEvent()
	_ = events.AddEvent(e)
	if err := events.DelEvent(e.Id); err != nil {
		t.Errorf("Error while delete event with ID = %d, error: %v", e.Id, err)
	}
	_, err := events.GetEvent(e.Id)
	if err == nil {
		t.Errorf("Error expected but was not returned: %v", err)
	}
}

func TestGetAllEvents(t *testing.T) {
	events := createNewDB()
	e1 := *calendar.NewEvent()
	_ = events.AddEvent(e1)
	e2 := *calendar.NewEvent()
	_ = events.AddEvent(e2)
	e3 := events.GetAllEvents()
	l := len(e3)
	if l != 2 {
		t.Errorf("After getting all events length slice != 2, actual length = %d", l)
		return
	}
}

func createNewDB() *calendar.DBMapEvents {
	return calendar.NewMapDB()
}
