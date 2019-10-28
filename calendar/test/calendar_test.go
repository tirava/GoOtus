/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 24.10.2019 21:15
 * Copyright (c) 2019 - Eugene Klimov
 */

package test

import (
	"github.com/evakom/calendar/internal/domain/interfaces"
	"github.com/evakom/calendar/internal/domain/models"
	"log"
	"os"
	"reflect"
	"testing"
)

const (
	EnvCalendarConfigPath  = "CALENDAR_CONFIG_PATH"
	FileCalendarConfigPath = "../internal/configs/calendar.yml"
)

func TestNewEvent(t *testing.T) {
	e1 := models.NewEvent().ID
	e2 := models.NewEvent().ID
	if e2 == e1 {
		t.Errorf("'id1 = %v' same as 'id2 = %v'", e1, e2)
	}
}

// will be later
//func TestProto(t *testing.T) {
//	e1 := &calendar.Event{}
//	e1.Subject = "testik"
//	out, err := proto.Marshal(e1)
//	if err != nil {
//		t.Fatalf("Failed to encode event: %v, error: %v", e1, err)
//	}
//	e2 := &calendar.Event{}
//	err = proto.Unmarshal(out, e2)
//	if err != nil {
//		t.Fatalf("Failed to parse event: %v, error: %v", out, err)
//	}
//	e2.XXX_sizecache = 8 // Subject length + 1
//	if !reflect.DeepEqual(e1, e2) {
//		t.Errorf("Event1 not equal Event2 after Marshal/Unmarshal:\n%#v\n%#v", e1, e2)
//	}
//}

func TestAddEvent(t *testing.T) {
	events := createNewDB()
	e := models.NewEvent()
	e.Subject = "222222222222222222222"
	e.Body = "3333333333333333333"
	_ = events.AddEvent(e)
	e = models.NewEvent()
	e.Duration = 555
	_ = events.AddEvent(e)
	l := len(events.GetAllEvents())
	if l != 2 {
		t.Errorf("After adding 2 events to MapDB length != 2, actual length = %d", l)
	}
}

func TestGetEvent(t *testing.T) {
	events := createNewDB()
	e1 := models.NewEvent()
	_ = events.AddEvent(e1)
	e2 := models.NewEvent()
	_ = events.AddEvent(e2)
	e3, _ := events.GetOneEvent(e1.ID)
	if !reflect.DeepEqual(e1, e3) {
		t.Errorf("Event1 not equal Event3 after get from DB:\n%#v\n%#v", e1, e3)
	}
}

func TestEditEvent(t *testing.T) {
	events := createNewDB()
	e1 := models.NewEvent()
	_ = events.AddEvent(e1)

	e2, _ := events.GetOneEvent(e1.ID)
	e2.Subject = "11111111111111111"
	e2.Body = "22222222222222222"
	_ = events.EditEvent(e2)

	e3, _ := events.GetOneEvent(e1.ID)
	if e3.Subject != e2.Subject || e3.Body != e2.Body {
		t.Errorf("Event1 not properly updated in the DB:\nExpected: %+v\nActual: %+v", e2, e3)
	}

	t2 := e2.UpdatedAt
	t3 := e3.UpdatedAt
	if t2.Equal(t3) {
		t.Errorf("Event1 updated time not correct in the DB:\nOld time: %v\nNew time: %v", t2, t3)
	}
}

func TestDelEvent(t *testing.T) {
	events := createNewDB()
	e := models.NewEvent()
	_ = events.AddEvent(e)
	if err := events.DelEvent(e.ID); err != nil {
		t.Errorf("Error while delete event with ID = %d, error: %v", e.ID, err)
	}
	_, err := events.GetOneEvent(e.ID)
	if err == nil {
		t.Errorf("Error expected but was not returned: %v", err)
	}
}

func TestGetAllEvents(t *testing.T) {
	events := createNewDB()
	e1 := models.NewEvent()
	_ = events.AddEvent(e1)
	e2 := models.NewEvent()
	_ = events.AddEvent(e2)
	e3 := events.GetAllEvents()
	l := len(e3)
	if l != 2 {
		t.Errorf("After getting all events length slice != 2, actual length = %d", l)
		return
	}
}

func createNewDB() interfaces.DB {
	confPath := os.Getenv(EnvCalendarConfigPath)
	if confPath == "" {
		confPath = FileCalendarConfigPath
	}
	conf := models.NewConfig(confPath)
	if err := conf.ReadParameters(); err != nil {
		log.Fatalln(err)
	}
	return interfaces.NewDB(conf.DBType)
}
