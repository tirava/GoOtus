# build
FROM golang:1.13.8-alpine as builder

WORKDIR /root
COPY . .
RUN CGO_ENABLED=0 go build -o app/http_server cmd/http_server/*.go

# run
FROM alpine:latest
LABEL maintainer="Eugene Klimov"

WORKDIR /root
COPY --from=builder /root/app/http_server .
EXPOSE 8000
CMD ["./http_server"]