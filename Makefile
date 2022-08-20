run-server:
	go run cmd/data_server/data_server.go

build:
	go build -o bin/data_server cmd/data_server/data_server.go

db-up:
	docker compose up
db-down:
	docker compose down

LOCAL_BIN:=$(CURDIR)/bin
.PHONY: .deps
.deps:
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go && \
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

MIGRATIONS_DIR=./migrations
.PHONY: migration
migration:
	goose -dir=${MIGRATIONS_DIR} create $(NAME) sql 

.PHONY: .test
.test:
	$(ino Running tests...)
	go test ./...

.PHONY: cover
cover:
	go test -v $$(go list ./... | grep -v -E './pkg/(api)') -covermode=count -coverprofile=/tmp/c.out
	go tool cover -html=/tmp/c.out