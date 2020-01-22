/*
 * Project: Image Previewer
 * Created on 22.01.2020 21:12
 * Copyright (c) 2020 - Eugene Klimov
 */

package http

import (
	"errors"
	"io"
	"net/http"

	"gitlab.com/tirava/image-previewer/internal/domain/preview"

	"gitlab.com/tirava/image-previewer/internal/models"
)

// Constants.
const (
	ReqIDField    = "request_id"
	HostField     = "host"
	MethodField   = "method"
	URLField      = "url"
	BrowserField  = "browser"
	RemoteField   = "remote"
	QueryField    = "query"
	CodeField     = "response_code"
	RespTimeField = "response_time"
)

type handler struct {
	handlers map[string]http.HandlerFunc
	preview  preview.Preview
	logger   models.Loggerer
	error    Error
}

func newHandlers(preview preview.Preview, logger models.Loggerer) *handler {
	return &handler{
		handlers: make(map[string]http.HandlerFunc),
		preview:  preview,
		logger:   logger,
		error:    newError(logger),
	}
}

func (h handler) helloHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")

	if name == "" {
		name = "nobody"
	}

	h.logger.WithFields(models.LoggerFields{
		CodeField:  http.StatusOK,
		ReqIDField: getRequestID(r.Context()),
	}).Infof("RESPONSE")

	s := "Hello, my name is " + name + "\n\n"

	if _, err := io.WriteString(w, s); err != nil {
		h.logger.Errorf("[hello] error write to response writer")
	}
}

func (h handler) previewHandler(w http.ResponseWriter, r *http.Request) {
	//key := urlform.FormEventID
	//value := r.URL.Query().Get(key)
	//if err := h.getEventsAndSend(key, value, w, r); err != nil {
	//	h.logger.Debugf("[getEvent] error: %s", err)
	//}
	h.error.send(w, 444, errors.New("555"), "666")
}
