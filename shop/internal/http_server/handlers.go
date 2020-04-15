package http

import (
	"net/http"
	"os"

	"github.com/go-chi/render"
)

func (s Server) root(w http.ResponseWriter, _ *http.Request) {
	host, err := os.Hostname()
	if err != nil {
		s.logger.Error().Msg(err.Error())
	}

	if _, err := w.Write([]byte("Welcome to our shop!\nMy pod hostname: " + host + "\n")); err != nil {
		s.logger.Error().Msg(err.Error())
	}
}

func (s Server) health(w http.ResponseWriter, r *http.Request) {
	render.Respond(w, r, render.M{"status": "OK"})
}
