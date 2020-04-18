package http

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog"
)

// Error model.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	logger  *zerolog.Logger
}

func newError(logger *zerolog.Logger) Error {
	return Error{
		logger: logger,
	}
}

func (e Error) send(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	errMsg := Error{
		Code:    code,
		Message: message,
	}

	if err := json.NewEncoder(w).Encode(errMsg); err != nil {
		e.logger.Error().Msg(err.Error())

		return
	}
}
