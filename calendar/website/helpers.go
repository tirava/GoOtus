/*
 * HomeWork-9: Calendar protobuf preparation
 * Created on 01.11.2019 13:17
 * Copyright (c) 2019 - Eugene Klimov
 */

package website

import (
	"context"
	"github.com/evakom/calendar/internal/loggers"
	"github.com/google/uuid"
	"net/http"
)

// Constants
const (
	IDField       = "id"
	HostField     = "host"
	MethodField   = "method"
	URLField      = "url"
	BrowserField  = "browser"
	RemoteField   = "remote"
	QueryField    = "query"
	CodeField     = "code"
	RespTimeField = "response_time"
)

type contextKey string

const contextKeyRequestID contextKey = "requestID"

func requestFields(r *http.Request, args ...string) loggers.Fields {
	fields := make(loggers.Fields)
	for _, s := range args {
		switch s {
		case IDField:
			fields[IDField] = getRequestID(r.Context())
		case HostField:
			fields[HostField] = r.Host
		case MethodField:
			fields[MethodField] = r.Method
		case URLField:
			fields[URLField] = r.URL.Path
		case BrowserField:
			fields[BrowserField] = r.Header.Get("User-Agent")
		case RemoteField:
			fields[RemoteField] = r.RemoteAddr
		case QueryField:
			fields[QueryField] = r.URL.RawQuery
		}
	}
	return fields
}

func assignRequestID(ctx context.Context) context.Context {
	reqID := uuid.New()
	return context.WithValue(ctx, contextKeyRequestID, reqID.String())
}

func getRequestID(ctx context.Context) string {
	reqID := ctx.Value(contextKeyRequestID)
	if key, ok := reqID.(string); ok {
		return key
	}
	return ""
}
