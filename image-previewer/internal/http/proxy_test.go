/*
 * Project: Image Previewer
 * Created on 24.01.2020 12:48
 * Copyright (c) 2020 - Eugene Klimov
 */

package http

import (
	"net/http"
	"reflect"
	"testing"
)

func TestProxyHeaders(t *testing.T) {
	fromClient := &http.Request{
		Header: make(http.Header),
	}

	fromClient.Header = http.Header{
		"X-111": {"111"},
		"X-222": {"222"},
		"X-333": {"333"},
	}

	noHeaders := []string{"x-222", "x-333"}
	expected := http.Header{
		"X-111": {"111"},
	}

	req, err := http.NewRequest("GET", "http://www.klim.go", nil)
	if err != nil {
		t.Fatal(err)
	}

	toServer := proxyHeaders(fromClient, req, noHeaders)

	if !reflect.DeepEqual(toServer.Header, expected) {
		t.Errorf("Expected headers:\n%s\n"+
			"but got headers:\n%s\n", expected, toServer.Header)
	}
}
