/*
 * Project: Image Previewer
 * Created on 22.01.2020 21:11
 * Copyright (c) 2020 - Eugene Klimov
 */

package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"gitlab.com/tirava/image-previewer/internal/models"
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
	logger         models.Loggerer
}

func newError(logger models.Loggerer) Error {
	return Error{
		logger: logger,
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
		e.logger.Errorf("can't encode error data: %s", err)
		return
	}
}
