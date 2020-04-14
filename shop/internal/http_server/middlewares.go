package http

import (
	"context"
	"net/http"
)

type stubs string

func stub(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), stubs("stub"), "stub")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
