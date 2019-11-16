/*
 * HomeWork-10: Calendar extending HTTP methods
 * Created on 16.11.2019 12:21
 * Copyright (c) 2019 - Eugene Klimov
 */

package website

import (
	"encoding/json"
	"errors"
	"github.com/evakom/calendar/internal/loggers"
	"net/http"
)

// ClientError model.
type ClientError struct {
	Error `json:"error"`
}

// Error model.
type Error struct {
	ErrCode        int    `json:"code"`
	ErrText        string `json:"text,omitempty"`
	ErrDescription string `json:"description"`
	logger         loggers.Logger
}

func newError() Error {
	return Error{
		logger: loggers.GetLogger(),
	}
}

func (e Error) send(w http.ResponseWriter, code int, err error, description string) {
	if err == nil {
		err = errors.New("")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	errMsg := ClientError{Error{
		ErrCode:        code,
		ErrText:        err.Error(),
		ErrDescription: description,
	}}
	if err := json.NewEncoder(w).Encode(errMsg); err != nil {
		e.logger.Error("can't encode error data: %s", err)
		return
	}
}
