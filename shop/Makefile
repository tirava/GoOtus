.PHONY: build run docker-build docker-run up down restart test

build:
	go build -o app/http_server cmd/http_server/*.go

run:
	./app/http_server

docker-build:
	docker build -t http_server -f ./build/package/http_server/Dockerfile .

docker-run:
	docker run --rm -d -p 8000:8000 --name http_server http_server app/http_server

up:
	docker-compose up -d --build

down:
	docker-compose down

restart: down up
