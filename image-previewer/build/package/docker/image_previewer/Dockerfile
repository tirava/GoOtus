FROM golang:1.13-alpine
LABEL maintainer="Eugene Klimov"

EXPOSE 8080
EXPOSE 9180

RUN apk add --no-cache make
RUN apk add --no-cache coreutils

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o app/image-previewer cmd/image-previewer/*.go
