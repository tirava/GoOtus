package http

import (
	"crypto/rand"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gitlab.com/tirava/image-previewer/internal/configs"
	"gitlab.com/tirava/image-previewer/internal/domain/entities"
	"gitlab.com/tirava/image-previewer/internal/domain/preview"
	"gitlab.com/tirava/image-previewer/internal/helpers"
	"gitlab.com/tirava/image-previewer/internal/loggers"
	"gitlab.com/tirava/image-previewer/internal/models"
)

const (
	fileConfigPath = "../../config.yml"
	imageWidth     = 500
	imageHeight    = 400
)

// nolint:gochecknoglobals
var testCases = []struct {
	description string
	urlPath     string
	imageName   string
	expectCode  int
	fromCache   string
}{
	{
		"correct request",
		"/preview/300/200/",
		"/image.jpg",
		http.StatusOK,
		"false",
	},
	{
		"incomplete parameters",
		"/preview",
		"",
		http.StatusBadRequest,
		"",
	},
	{
		"bad preview size",
		"/preview/qqq/www/",
		"",
		http.StatusBadRequest,
		"",
	},
	{
		"image name without ext not found",
		"/preview/300/200/",
		"/imagejpg",
		http.StatusNotFound,
		"false",
	},
	{
		"bad image type",
		"/preview/300/200/",
		"/image.tiff",
		http.StatusInternalServerError,
		"false",
	},
	{
		"correct request from cache",
		"/preview/300/200/",
		"/image.jpg",
		http.StatusOK,
		"true",
	},
}

func initConfLogger() (models.Loggerer, *preview.Preview) {
	cfg, err := configs.NewConfig(fileConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	conf := cfg.GetConfig()
	lg, err := loggers.NewLogger(conf.Logger, "none", ioutil.Discard)

	if err != nil {
		log.Fatal(err)
	}

	conf.Cacher = "nolimit"
	conf.Storager = "inmemory"
	conf.StoragePath = ""
	conf.MaxCacheItems = 0
	prev, err := helpers.InitPreview(conf)

	if err != nil {
		log.Fatal(err)
	}

	return lg, prev
}

func TestGetHello(t *testing.T) {
	var handlers *handler

	lg, prev := initConfLogger()

	handlers = newHandlers(lg, models.Config{}, *prev, entities.ResizeOptions{})

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	query := req.URL.Query()
	//query.Add("name", "Klim")
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

	expected := "Hello, my name is Previewer!"
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("Hello handler returned unexpected body:\ngot - %v\nwant - %v",
			rr.Body.String(), expected)
		return
	}
}

func TestPreviewHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/image.jpg" && r.URL.Path != "/image.tiff" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if r.URL.Path == "/image.tiff" {
			buf := make([]byte, 4)
			_, _ = rand.Read(buf)
			_, _ = w.Write(buf)
			return
		}
		im := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
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

	handlers = newHandlers(lg, models.Config{}, *prev, entities.ResizeOptions{})
	handler := http.HandlerFunc(handlers.previewHandler)

	for _, test := range testCases {
		req := httptest.NewRequest(http.MethodGet, test.urlPath+ts.URL+test.imageName, nil)
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != test.expectCode {
			t.Errorf("Test: %s\n"+
				"Preview handler returned wrong status code: got - %v, want - %v\n"+
				"response: %s",
				test.description, status, test.expectCode, rr.Body)
		}

		if cached := rr.Header().Get("From-Cache"); cached != test.fromCache {
			t.Errorf("Test: %s\n"+
				"Preview handler returned wrong cache status: got - %v, want - %v\n",
				test.description, cached, test.fromCache)
		}
	}
}
