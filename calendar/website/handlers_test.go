/*
 * HomeWork-9: Calendar protobuf preparation
 * Created on 03.11.2019 13:01
 * Copyright (c) 2019 - Eugene Klimov
 */

package website

import (
	"github.com/evakom/calendar/internal/domain/calendar"
	"github.com/evakom/calendar/internal/loggers"
	"github.com/evakom/calendar/tools"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const fileConfigPath = "../config.yml"

var handlers *handler

func init() {
	conf := tools.InitConfig(fileConfigPath)
	loggers.NewLogger("none", nil)
	db := tools.InitDB(conf.DBType)
	cal := calendar.NewCalendar(db)
	handlers = newHandlers(cal)
}

func TestGetHello(t *testing.T) {

	req := httptest.NewRequest("GET", "/", nil)

	query := req.URL.Query()
	query.Add("name", "Klim")
	query.Add("qqq", "www") // fake
	req.URL.RawQuery = query.Encode()

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.helloHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Hello handler returned wrong status code: got - %v, want - %v",
			status, http.StatusOK)
		return
	}

	expected := "Hello, my name is Klim"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Hello handler returned unexpected body:\ngot - %v\nwant - %v",
			rr.Body.String(), expected)
		return
	}
}
