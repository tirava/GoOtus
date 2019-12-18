FROM alpine:latest
LABEL maintainer="Eugene Klimov"

RUN apk add --no-cache coreutils
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

WORKDIR /app
COPY ./integration_tests/integration_tests /app
COPY ./integration_tests/config.yml /app
COPY ./integration_tests/features/. /app/features

ENTRYPOINT /app/integration_tests
