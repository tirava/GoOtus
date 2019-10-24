/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 24.10.2019 21:15
 * Copyright (c) 2019 - Eugene Klimov
 */

package calendar

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"reflect"
	"strings"
	"testing"
)

func TestNewEvent(t *testing.T) {
	e1 := newEvent().GetId()
	e2 := newEvent().GetId()
	if e2 == e1 {
		t.Errorf("'id1 = %v' same as 'id2 = %v'", e1, e2)
	}
}

func TestNewDB(t *testing.T) {
	if _, err := createNewDB(); err != nil {
		t.Fatalf("Can't create new DB, error: %v", err)
	}
}

func TestProto(t *testing.T) {
	e1 := &Event{}
	e1.Subject = "testik"
	out, err := proto.Marshal(e1)
	if err != nil {
		t.Fatalf("Failed to encode event: %v, error: %v", e1, err)
	}
	e2 := &Event{}
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
	events, _ := createNewDB()
	e := *newEvent()
	e.Subject = "222222222222222222222"
	e.Body = "3333333333333333333"
	_ = events.addEvent(e)
	e = *newEvent()
	e.Duration = 555
	_ = events.addEvent(e)
	l := len(events.events)
	if l != 2 {
		t.Errorf("After adding 2 events to MapDB length != 2, actual length = %d", l)
	}
}

func TestEditEvent(t *testing.T) {
	events, _ := createNewDB()
	e1 := *newEvent()
	_ = events.addEvent(e1)

	e2, _ := events.getEvent(e1.Id)
	e2.Subject = "11111111111111111"
	e2.Body = "22222222222222222"
	_ = events.editEvent(e2)

	e3, _ := events.getEvent(e1.Id)
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
	events, _ := createNewDB()
	e := *newEvent()
	_ = events.addEvent(e)
	if err := events.delEvent(e.Id); err != nil {
		t.Errorf("Error while delete event with ID = %d, error: %v", e.Id, err)
	}
	_, err := events.getEvent(e.Id)
	if err == nil {
		t.Errorf("Error expected but was not returned: %v", err)
	}
}

func TestGetEvent(t *testing.T) {

}

func TestGetAllEvents(t *testing.T) {

}

func createNewDB() (*dbMapEvents, error) {
	events := newDB(&dbMapEvents{})
	if db, ok := events.(*dbMapEvents); !ok {
		return nil, fmt.Errorf("error while cast interface{} to *dbMapEvents:\n%#v", events)
	} else {
		return db, nil
	}
}
