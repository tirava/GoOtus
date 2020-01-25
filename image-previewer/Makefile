.PHONY: build run docker-build docker-run up down restart

build:
	go build -o app/image-previewer cmd/image-previewer/*.go
	go build -o app/image-previewer-cli cmd/image-previewer-cli/*.go

run: build
	./app/image-previewer

docker-build:
	docker build -t image-previewer-http-server -f ./deployments/docker/image_previewer/Dockerfile .

docker-run: docker-build
	docker run --rm -d -p 8080:8080 -p 9180:9180 --name image-previewer-http-server image-previewer-http-server app/image-previewer

up:
	docker-compose up -d --build

down:
	docker-compose down

restart: down up

