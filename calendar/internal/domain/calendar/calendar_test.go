/*
 * HomeWork-9: Calendar protobuf preparation
 * Created on 04.11.2019 18:02
 * Copyright (c) 2019 - Eugene Klimov
 */

package calendar

import (
	"github.com/evakom/calendar/internal/domain/errors"
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/evakom/calendar/tools"
	"github.com/google/uuid"
	"testing"
)

const fileConfigPath = "../../../config.yml"

var cal Calendar

func init() {
	conf := tools.InitConfig(fileConfigPath)
	models.NewLogger("none", nil)
	//models.NewLogger("debug", os.Stdout)
	db := tools.InitDB(conf.DBType)
	cal = NewCalendar(db)
}

func TestAddEvent(t *testing.T) {
	e := models.NewEvent()
	e.Subject = "44444444444444444"
	e.Body = "55555555555555555"
	if err := cal.AddEvent(e); err != nil {
		t.Errorf("Adding calendar event should not return error but returns it: %s", err)
	}
	e.Duration = 666
	if err := cal.AddEvent(e); err == nil {
		t.Errorf("Adding event with same id should return error but returns no error")
	}
}

func TestGetAllEventsFilterEventID(t *testing.T) {
	e1 := models.NewEvent()
	_ = cal.AddEvent(e1)
	e2 := models.NewEvent()
	_ = cal.AddEvent(e2)
	filter := models.Event{ID: e1.ID}
	events, err := cal.GetAllEventsFilter(filter)
	if err != nil {
		t.Errorf("Filtered event by id should not return error but returns it: %s", err)
		return
	}
	if events[0] != e1 {
		t.Errorf("Added event with 'id=%v' but filtered with 'id=%v'", e1.ID, events[0].ID)
	}
	filter = models.Event{ID: uuid.New()}
	_, err = cal.GetAllEventsFilter(filter)
	if err != errors.ErrEventNotFound {
		t.Errorf("Returned error is not expected type, actual: %s, "+
			"but expected: %s", err, errors.ErrEventNotFound)
	}
}

func TestGetAllEventsFilterUserID(t *testing.T) {
	e1 := models.NewEvent()
	_ = cal.AddEvent(e1)
	e2 := models.NewEvent()
	_ = cal.AddEvent(e2)
	e3 := models.NewEvent()
	e3.UserID = e1.UserID
	_ = cal.AddEvent(e3)
	filter := models.Event{UserID: e1.UserID}
	events, err := cal.GetAllEventsFilter(filter)
	if err != nil {
		t.Errorf("Filtered event by user should not return error but returns it: %s", err)
		return
	}
	l := len(events)
	if l != 2 {
		t.Errorf("After getting filtered events length slice != 2, actual length = %d", l)
	}
	for _, e := range events {
		if e.UserID != filter.UserID {
			t.Errorf("UserID=%s in event id=%s not equal\n"+
				"expected userID=%s", e.UserID.String(), e.ID.String(), filter.UserID)
		}
	}
	filter = models.Event{}
	events, err = cal.GetAllEventsFilter(filter)
	if events != nil || err != nil {
		t.Errorf("Null filter must return nil events but returned:\n"+
			"events=%v, err=%s", events, err)
	}
}
