gen:
	protoc --go_out=plugins=grpc:internal/grpc api/*.proto

build: gen
	go build -o http_server cmd/http_server/main.go
	go build -o grpc_server cmd/grpc_server/main.go
	go build -o alert_scheduler cmd/notifications/scheduler/*.go
	go build -o alert_sender cmd/notifications/sender/*.go
	go build -o grpc_client cmd/grpc_client/main.go

test: build
	go test ./...

docker-build: build
	docker build -t calendar-http-server -f ./deployments/docker/http_server/Dockerfile .
	docker build -t calendar-grpc-server -f ./deployments/docker/grpc_server/Dockerfile .
	docker build -t calendar-alert-scheduler -f ./deployments/docker/alert_scheduler/Dockerfile .
	docker build -t calendar-alert-sender -f ./deployments/docker/alert_sender/Dockerfile .

docker-run: docker-build
	docker run --rm -d -p 8080:8080 --name calendar-http-server --network calendar_calendar calendar-http-server
	docker run --rm -d -p 8888:8888 --name calendar-grpc-server --network calendar_calendar calendar-grpc-server
	docker run --rm -d --name calendar-alert-scheduler --network calendar_calendar calendar-alert-scheduler
	docker run --rm -d --name calendar-alert-sender --network calendar_calendar calendar-alert-sender