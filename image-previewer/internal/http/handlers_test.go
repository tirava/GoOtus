/*
 * HomeWork-9: Calendar protobuf preparation
 * Created on 03.11.2019 13:01
 * Copyright (c) 2019 - Eugene Klimov
 */

package http

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"

	"gitlab.com/tirava/image-previewer/internal/configs"
	"gitlab.com/tirava/image-previewer/internal/domain/preview"
	"gitlab.com/tirava/image-previewer/internal/loggers"
)

const fileConfigPath = "../../config.yml"

func TestGetHello(t *testing.T) {
	var handlers *handler

	cfg, err := configs.NewConfig(fileConfigPath)
	if err != nil {
		t.Fatal(err)
	}

	conf := cfg.GetConfig()
	lg, err := loggers.NewLogger(conf.Logger, "none", ioutil.Discard)

	if err != nil {
		t.Fatal(err)
	}

	prev, err := preview.NewPreview(conf.Previewer)
	if err != nil {
		t.Fatal(err)
	}

	handlers = newHandlers(lg, prev, entities.ResizeOptions{})

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
