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
	error         Error
}

func newHandlers(logger models.Loggerer,
	preview preview.Preview, opts entities.ResizeOptions) *handler {
	return &handler{
		handlers:      make(map[string]http.HandlerFunc),
		cacheHandlers: make(map[string]*regexp.Regexp),
		preview:       preview,
		opts:          opts,
		logger:        logger,
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
	extDot := strings.LastIndex(source, ".")

	if extDot == -1 {
		errSend := errors.New("unknown image extension")
		h.errorHelper(w, r, http.StatusBadRequest, errors.New(""), errSend, "unknown image extension %s",
			"make sure the request image with extension (jpeg, jpg or png)")

		return
	}

	ext := source[extDot:]
	img, err := getSourceImage(source, ext)

	if err != nil {
		errSend := fmt.Errorf("invalid image source: %w", err)
		h.errorHelper(w, r, http.StatusNoContent, err, errSend, "invalid image source: %s",
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
	case ".jpg", ".jpeg":
		contType = "image/jpeg"
		err = jpeg.Encode(buffer, img, nil)
	case ".png":
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

func getSourceImage(url, imgType string) (image.Image, error) {
	var err error

	var img image.Image

	prefix := ""
	if !strings.HasPrefix(url, "http") {
		prefix = "https://"
	}

	response, err := http.Get(fmt.Sprintf("%s%s", prefix, url))
	if err != nil {
		return nil, fmt.Errorf("unable to get image from source: %w", err)
	}
	defer response.Body.Close()

	switch imgType {
	case ".jpg", ".jpeg":
		img, err = jpeg.Decode(response.Body)
	case ".png":
		img, err = png.Decode(response.Body)
	default:
		return img, errors.New("bad image extension")
	}

	if err != nil {
		return img, fmt.Errorf("unable to decode image: %w", err)
	}

	return img, nil
}
