/*
 * Project: Image Previewer
 * Created on 30.01.2020 16:22
 * Copyright (c) 2020 - Eugene Klimov
 */

package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/DATA-DOG/godog"
)

type previewTest struct {
	responseStatusCode int
	responseBody       []byte
	data               string
	cacheHeaderVal     string
}

func (t *previewTest) iSendRequestTo(httpMethod, addr string) error {
	var r *http.Response

	var err error

	switch httpMethod {
	case http.MethodGet:
		// nolint
		r, err = http.Get(addr)
	default:
		err = fmt.Errorf("unknown method: %s", httpMethod)
	}

	if err != nil {
		return err
	}

	t.responseStatusCode = r.StatusCode
	t.responseBody, err = ioutil.ReadAll(r.Body)

	if err != nil {
		return err
	}
	defer r.Body.Close()

	t.cacheHeaderVal = r.Header.Get("From-Cache")

	return nil
}

func (t *previewTest) theResponseCodeShouldBe(code int) error {
	if t.responseStatusCode != code {
		return fmt.Errorf("unexpected status code: %d != %d", t.responseStatusCode, code)
	}

	return nil
}

func (t *previewTest) theResponseShouldMatchText(text string) error {
	if string(t.responseBody) != text {
		fmt.Println([]byte(text), t.responseBody)
		return fmt.Errorf("unexpected text: %s != %s", t.responseBody, text)
	}

	return nil
}

func (t *previewTest) iReceiveErrorWithText(text string) error {
	if !strings.Contains(string(t.responseBody), text) {
		fmt.Println([]byte(text), t.responseBody)
		return fmt.Errorf("unexpected text: %s != %s", t.responseBody, text)
	}

	return nil
}

func (t *previewTest) iSendRequestToWithData(httpMethod, addr, data string) error {
	t.data = data

	if err := t.iSendRequestTo(httpMethod, addr+data); err != nil {
		return err
	}

	return nil
}

func (t *previewTest) iReceiveImageWithSize(size string) error {
	sXY := strings.Split(size, "x")
	x, errX := strconv.Atoi(sXY[0])
	y, errY := strconv.Atoi(sXY[1])

	if errX != nil || errY != nil {
		return fmt.Errorf("error conversion from image size: x='%s', y='%s'", sXY[0], sXY[1])
	}

	br := bytes.NewReader(t.responseBody)
	img, err := jpeg.Decode(br)

	if err != nil {
		return fmt.Errorf("unable to decode image: %s", err)
	}

	if img.Bounds().Max.X != x || img.Bounds().Max.Y != y {
		return fmt.Errorf("unexpected preview image size: %d != %d, %d != %d",
			x, img.Bounds().Max.X, y, img.Bounds().Max.Y)
	}

	return nil
}

func (t *previewTest) iReceivedHeader(key, value string) error {
	if t.cacheHeaderVal != value {
		return fmt.Errorf("unexpected cache header '%s': %s != %s", key, t.cacheHeaderVal, value)
	}

	return nil
}

func FeatureContext(s *godog.Suite) {
	test := new(previewTest)

	s.Step(`^I send "([^"]*)" request to "([^"]*)"$`, test.iSendRequestTo)
	s.Step(`^The response code should be (\d+)$`, test.theResponseCodeShouldBe)
	s.Step(`^The response should match text "([^"]*)"$`, test.theResponseShouldMatchText)
	s.Step(`^I receive error with text "([^"]*)"$`, test.iReceiveErrorWithText)
	s.Step(`^I send "([^"]*)" request to "([^"]*)" with data "([^"]*)"$`, test.iSendRequestToWithData)
	s.Step(`^I receive image with size "([^"]*)"$`, test.iReceiveImageWithSize)
	s.Step(`^I received header "([^"]*)" = "([^"]*)"$`, test.iReceivedHeader)
}
