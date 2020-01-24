/*
 * Project: Image Previewer
 * Created on 22.01.2020 21:12
 * Copyright (c) 2020 - Eugene Klimov
 */

package http

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"

	"gitlab.com/tirava/image-previewer/internal/domain/preview"

	"gitlab.com/tirava/image-previewer/internal/models"
)

// Constants.
const (
	ReqIDField          = "request_id"
	HostField           = "host"
	MethodField         = "method"
	URLField            = "url"
	BrowserField        = "browser"
	RemoteField         = "remote"
	QueryField          = "query"
	CodeField           = "response_code"
	RespTimeField       = "response_time"
	minRequestLenParams = 5
)

type handler struct {
	handlers      map[string]http.HandlerFunc
	cacheHandlers map[string]*regexp.Regexp
	preview       preview.Preview
	opts          entities.ResizeOptions
	logger        models.Loggerer
	noHeaders     []string
	error         Error
}

func newHandlers(logger models.Loggerer, noHeaders []string,
	preview preview.Preview, opts entities.ResizeOptions) *handler {
	return &handler{
		handlers:      make(map[string]http.HandlerFunc),
		cacheHandlers: make(map[string]*regexp.Regexp),
		preview:       preview,
		opts:          opts,
		logger:        logger,
		noHeaders:     noHeaders,
		error:         newError(logger),
	}
}

func (h handler) errorHelper(w http.ResponseWriter, r *http.Request, code int,
	err, errSend error, message, description string) {
	h.logger.WithFields(models.LoggerFields{
		CodeField:  code,
		ReqIDField: getRequestID(r.Context()),
	}).Errorf(message, err)
	h.error.send(w, code, errSend, description)
}

func (h handler) helloHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")

	if name == "" {
		name = "nobody"
	}

	h.logger.WithFields(models.LoggerFields{
		CodeField:  http.StatusOK,
		ReqIDField: getRequestID(r.Context()),
	}).Infof("RESPONSE")

	s := "Hello, my name is " + name + "\n\n"

	if _, err := io.WriteString(w, s); err != nil {
		h.logger.Errorf("[hello] error write to response writer")
	}
}

func (h handler) previewHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.SplitN(r.URL.Path, "/", minRequestLenParams)
	if len(path) < minRequestLenParams {
		errSend := fmt.Errorf("invalid request params: %s", r.URL.Path)
		h.errorHelper(w, r, http.StatusBadRequest, errors.New(""), errSend, "invalid request params %s",
			"make sure the request contains 'preview/sizeX/sizeY/source_image'")

		return
	}

	sX, sY := path[2], path[3]
	previewX, errX := strconv.Atoi(sX)
	previewY, errY := strconv.Atoi(sY)

	if errX != nil || errY != nil {
		errSend := fmt.Errorf("invalid request preview size: %sx%s", sX, sY)
		h.errorHelper(w, r, http.StatusBadRequest, errors.New(""), errSend, "invalid request preview size %s",
			"make sure the request contains integer preview size")

		return
	}

	source := path[4]

	img, ext, err := h.getSourceImage(r, source)

	if err != nil {
		errSend := fmt.Errorf("invalid image source: %w", err)
		h.errorHelper(w, r, http.StatusNotFound, err, errSend, "invalid image source: %s",
			"make sure the request contains valid image source url")

		return
	}

	previewed := h.preview.Preview(previewX, previewY, img, h.opts)

	if err := h.writeImage(w, previewed, ext); err != nil {
		errSend := fmt.Errorf("invalid image type: %w", err)
		h.errorHelper(w, r, http.StatusInternalServerError, err, errSend, "invalid image type: %s",
			"make sure the request valid image type (jpeg or png)")

		return
	}

	h.logger.WithFields(models.LoggerFields{
		CodeField:  http.StatusOK,
		ReqIDField: getRequestID(r.Context()),
	}).Infof("previewed successfully")
}

func (h handler) writeImage(w http.ResponseWriter, img image.Image, imgType string) error {
	var err error

	buffer := new(bytes.Buffer)
	contType := ""

	switch imgType {
	case "jpg", "jpeg":
		contType = "image/jpeg"
		err = jpeg.Encode(buffer, img, nil)
	case "png":
		contType = "image/png"
		err = png.Encode(buffer, img)
	default:
		return errors.New(imgType)
	}

	if err != nil {
		return errors.New("unable to encode image")
	}

	w.Header().Set("Content-Type", contType)
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))

	if _, err := w.Write(buffer.Bytes()); err != nil {
		return errors.New("unable to write image to response writer")
	}

	return nil
}

func (h handler) getSourceImage(r *http.Request, url string) (image.Image, string, error) {
	var err error

	var img image.Image

	prefix := ""
	if !strings.HasPrefix(url, "http") {
		prefix = "https://"
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", prefix, url), nil)

	if err != nil {
		return img, "", fmt.Errorf("unable create new request: %w", err)
	}

	response, err := client.Do(proxyHeaders(r, req, h.noHeaders))

	if err != nil {
		return img, "", fmt.Errorf("unable to get image from source: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return img, "", fmt.Errorf("remote server returns bad status: %d", response.StatusCode)
	}

	img, ext, err := image.Decode(response.Body)

	if err != nil {
		return img, ext, fmt.Errorf("unable to decode image: %w", err)
	}

	return img, ext, nil
}
