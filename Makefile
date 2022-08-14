.PHONY: run
run-bot:
	go run cmd/bot/main.go

run-server:
	go run cmd/server/server.go

build:
	go build -o bin/bot cmd/bot/bot.go
	go build -o bin/server cmd/server/server.go

LOCAL_BIN:=$(CURDIR)/bin
.PHONY: .deps
.deps:
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
