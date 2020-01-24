/*
 * Project: Image Previewer
 * Created on 23.01.2020 13:20
 * Copyright (c) 2020 - Eugene Klimov
 */

package http

import (
	"image"
	"image/gif"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gitlab.com/tirava/image-previewer/internal/models"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"

	"gitlab.com/tirava/image-previewer/internal/configs"
	"gitlab.com/tirava/image-previewer/internal/domain/preview"
	"gitlab.com/tirava/image-previewer/internal/loggers"
)

const (
	fileConfigPath = "../../config.yml"
	imageWidth     = 500
	imageHeight    = 400
)

// nolint
var testCases = []struct {
	description string
	urlPath     string
	imageName   string
	expectCode  int
}{
	{
		"correct request",
		"/preview/300/200/",
		"/image.jpg",
		http.StatusOK,
	},
	{
		"incomplete parameters",
		"/preview",
		"",
		http.StatusBadRequest,
	},
	{
		"bad preview size",
		"/preview/qqq/www/",
		"",
		http.StatusBadRequest,
	},
	{
		"image name without ext not found",
		"/preview/300/200/",
		"/imagejpg",
		http.StatusNotFound,
	},
	{
		"bad image type",
		"/preview/300/200/",
		"/image.gif",
		http.StatusInternalServerError,
	},
}

func initConfLogger() (models.Loggerer, preview.Preview) {
	cfg, err := configs.NewConfig(fileConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	conf := cfg.GetConfig()
	lg, err := loggers.NewLogger(conf.Logger, "none", ioutil.Discard)

	if err != nil {
		log.Fatal(err)
	}

	prev, err := preview.NewPreview(conf.Previewer)
	if err != nil {
		log.Fatal(err)
	}

	return lg, prev
}

func TestGetHello(t *testing.T) {
	var handlers *handler

	lg, prev := initConfLogger()

	handlers = newHandlers(lg, []string{}, prev, entities.ResizeOptions{})

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

func TestPreviewHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/image.jpg" && r.URL.Path != "/image.gif" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		im := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
		if r.URL.Path == "/image.gif" {
			_ = gif.Encode(w, im, nil)
			return
		}
		_ = jpeg.Encode(w, im, nil)
	}))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	var handlers *handler

	lg, prev := initConfLogger()

	handlers = newHandlers(lg, []string{}, prev, entities.ResizeOptions{})
	handler := http.HandlerFunc(handlers.previewHandler)

	for _, test := range testCases {
		req := httptest.NewRequest("GET", test.urlPath+ts.URL+test.imageName, nil)
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != test.expectCode {
			t.Errorf("Test: %s\n"+
				"Preview handler returned wrong status code: got - %v, want - %v\n"+
				"response: %s",
				test.description, status, test.expectCode, rr.Body)
		}
	}
}
