/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 24.10.2019 21:15
 * Copyright (c) 2019 - Eugene Klimov
 */

package dbs

import (
	"github.com/evakom/calendar/internal/configs"
	"github.com/evakom/calendar/internal/domain/interfaces"
	"github.com/evakom/calendar/internal/domain/models"
	uuid "github.com/satori/go.uuid"
	"log"
	"os"
	"reflect"
	"testing"
)

const (
	EnvCalendarConfigPath  = "CALENDAR_CONFIG_PATH"
	FileCalendarConfigPath = "../../config.yml"
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
	if err := events.AddEvent(e); err != nil {
		t.Errorf("Adding event should not return error but returns it: %s", err)
	}
	e = models.NewEvent()
	e.Duration = 555
	_ = events.AddEvent(e)
	l := len(events.GetAllEvents())
	if l != 2 {
		t.Errorf("After adding 2 events to MapDB length != 2, actual length = %d", l)
	}
	if err := events.AddEvent(e); err == nil {
		t.Errorf("Adding event with same id should return error but returns no error")
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
	e3.ID = uuid.NewV4()
	if _, err := events.GetOneEvent(e3.ID); err == nil {
		t.Errorf("Error expected but was not returned for getting id = : %s", e3.ID.String())
	}
}

func TestEditEvent(t *testing.T) {
	events := createNewDB()
	e1 := models.NewEvent()
	_ = events.AddEvent(e1)

	e2, _ := events.GetOneEvent(e1.ID)
	e2.Subject = "11111111111111111"
	e2.Body = "22222222222222222"
	if err := events.EditEvent(e2); err != nil {
		t.Errorf("Editing event should not return error but returns it: %s", err)
	}

	e3, _ := events.GetOneEvent(e1.ID)
	if e3.Subject != e2.Subject || e3.Body != e2.Body {
		t.Errorf("Event1 not properly updated in the DB:\nExpected: %+v\nActual: %+v", e2, e3)
	}

	t2 := e2.UpdatedAt
	t3 := e3.UpdatedAt
	if t2.Equal(t3) {
		t.Errorf("Event1 updated time not correct in the DB:\nOld time: %v\nNew time: %v", t2, t3)
	}
	e3.ID = uuid.NewV4()
	if err := events.EditEvent(e3); err == nil {
		t.Errorf("Editing event with same id should return error but returns no error")
	}
}

func TestDelEvent(t *testing.T) {
	events := createNewDB()
	e := models.NewEvent()
	_ = events.AddEvent(e)
	if err := events.DelEvent(e.ID); err != nil {
		t.Errorf("Error while delete event with ID = %d, error: %v", e.ID, err)
	}
	if _, err := events.GetOneEvent(e.ID); err == nil {
		t.Errorf("Error expected but was not returned for deleted id = : %s", e.ID.String())
	}
	e.ID = uuid.NewV4()
	if err := events.DelEvent(e.ID); err == nil {
		t.Errorf("Error expected but was not returned for deleting fake id = %s", e.ID.String())
	}
}

func TestGetAllEvents(t *testing.T) {
	events := createNewDB()
	e1 := models.NewEvent()
	_ = events.AddEvent(e1)
	e2 := models.NewEvent()
	_ = events.AddEvent(e2)
	_ = events.DelEvent(e1.ID)
	e3 := events.GetAllEvents()
	l := len(e3)
	if l != 1 {
		t.Errorf("After getting all events length slice != 1, actual length = %d", l)
	}
}

func createNewDB() interfaces.DB {
	confPath := os.Getenv(EnvCalendarConfigPath)
	if confPath == "" {
		confPath = FileCalendarConfigPath
	}
	conf := configs.NewConfig(confPath)
	if err := conf.ReadParameters(); err != nil {
		log.Fatalln(err)
	}
	models.NewLogger("none", nil)
	db, err := NewDB(conf.DBType)
	if db == nil {
		log.Fatalf("unsupported DB type: %s\n", conf.DBType)
	}
	if err != nil {
		log.Fatalf("Open DB: %s, error: %s \n", conf.DBType, err)
	}
	return db
}
