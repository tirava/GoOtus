.PHONY: build run docker-build docker-run docker-stop

build:
	go build -o app/http_server cmd/http_server/*.go

run:
	./app/http_server

docker-build:
	docker build -t http_server -f ./build/package/http_server/Dockerfile .

docker-run:
	docker run --rm -d -p 8000:8000 --name http_server http_server

docker-stop:
	docker stop http_server
