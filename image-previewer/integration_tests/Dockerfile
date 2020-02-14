FROM golang:1.13-alpine
LABEL maintainer="Eugene Klimov"

ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR /opt/integration_tests

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go test ./integration_tests/... -o ./integration_tests/integration_tests -i

WORKDIR /opt/integration_tests/integration_tests