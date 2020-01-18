.PHONY: build run docker-build docker-run up down restart test

build:
	go build -o app/image-previewer cmd/image-previewer/*.go
#	go build -o grpc_server cmd/grpc_server/main.go

run: build
	./app/image-previewer

docker-build:
	docker build -t image-previewer-http-server -f ./deployments/docker/image_previewer/Dockerfile .
# 	docker build -t calendar-grpc-server -f ./deployments/docker/grpc_server/Dockerfile .

docker-run: docker-build
	docker run --rm -d -p 8080:8080 --name image-previewer-http-server image-previewer-http-server
# 	docker run --rm -d -p 8888:8888 --name calendar-grpc-server --network calendar_calendar calendar-grpc-server

up:
	docker-compose up -d --build

down:
	docker-compose down

restart: down up

# test:
# 	set -e ;\
# 	go test ./integration_tests/... -o ./integration_tests/integration_tests -i ;\
#     docker-compose -f docker-compose.test.yml build ;\
#     test_status_code=0 ;\
#     docker-compose -f docker-compose.test.yml run calendar-integration-tests || test_status_code=$$? ;\
#     docker-compose -f docker-compose.test.yml down ;\
#     exit $$test_status_code ;\