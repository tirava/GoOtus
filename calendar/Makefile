gen:
	protoc --go_out=plugins=grpc:internal/grpc api/*.proto

build: gen
	go build -o grpc_client cmd/grpc_client/main.go
	go build -o grpc_server cmd/grpc_server/main.go
	go build -o http_server cmd/http_server/main.go

test: build
	go test ./...
