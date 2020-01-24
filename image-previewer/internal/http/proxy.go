/*
 * Project: Image Previewer
 * Created on 23.01.2020 23:42
 * Copyright (c) 2020 - Eugene Klimov
 */

package http

import (
	"net/http"
	"sort"
	"strings"
)

// proxyHeaders - noHeaders must be sorted ascending and in lowercase
func proxyHeaders(fromClient, toServer *http.Request, noHeaders []string) *http.Request {
	for k, v := range fromClient.Header {
		ks := strings.ToLower(k)
		i := sort.SearchStrings(noHeaders, ks)

		if i != len(noHeaders) && noHeaders[i] == ks {
			continue
		}

		vs := make([]string, len(v))
		copy(vs, v)
		toServer.Header[k] = vs
	}

	return toServer
}
