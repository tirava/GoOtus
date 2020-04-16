.PHONY: build run docker-build docker-run docker-stop

build:
	go build -o app/http_server cmd/http_server/*.go

run:
	./app/http_server

docker-build:
	docker build -t evaklim/shop -f ./build/package/http_server/Dockerfile .

docker-run:
	docker run --rm -d -p 8000:8000 --name evaklim_shop evaklim/shop

docker-stop:
	docker stop evaklim_shop
