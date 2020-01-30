.PHONY: build run docker-build docker-run up down restart test

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

test:
	set -e ;\
    docker-compose -f docker-compose.test.yml up --build -d ;\
    test_status_code=0 ;\
    docker-compose -f docker-compose.test.yml run image-previewer-integration-tests ;\
    go test ./integration_tests/... -v || test_status_code=$$? ;\
    docker-compose -f docker-compose.test.yml down ;\
    exit $$test_status_code ;\