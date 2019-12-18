FROM alpine:latest
LABEL maintainer="Eugene Klimov"

RUN apk add --no-cache coreutils
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

EXPOSE 8080

WORKDIR /app
COPY http_server /app
COPY config.yml /app

ENTRYPOINT /app/http_server
