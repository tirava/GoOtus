gen:
	protoc --go_out=plugins=grpc:internal/grpc api/*.proto

build: gen
	go build -o calendar main.go