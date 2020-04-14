package http

import (
	"net/http"

	"github.com/go-chi/render"
)

func (s Server) root(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("Welcome to our shop!")); err != nil {
		//log.
	}
}

func (s Server) health(w http.ResponseWriter, r *http.Request) {
	render.Respond(w, r, render.M{"status": "OK"})
}
