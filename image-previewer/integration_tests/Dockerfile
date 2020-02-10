FROM golang:1.13-alpine
LABEL maintainer="Eugene Klimov"

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

WORKDIR /opt/integration_tests

COPY . .