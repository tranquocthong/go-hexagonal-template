SVC=yourservice
BIN=bin/$(SVC)

SWAG_BIN=$(shell go env GOPATH)/bin/swag

LINT_BIN=$(shell which golangci-lint)

.PHONY: build run tidy test lint fmt docker docker-run swagger

swagger:
	$(SWAG_BIN) init -g cmd/service/main.go -o ./docs/swagger

build:
	@mkdir -p bin
	GO111MODULE=on CGO_ENABLED=0 go build -o $(BIN) ./cmd/service

run: swagger build
	APP_NAME=$(SVC) $(BIN)

lint:
	$(LINT_BIN) run ./...

tidy:
	go mod tidy

test:
	go test ./...

fmt:
	go fmt ./...

docker:
	docker build -t $(SVC):local .

docker-run:
	docker run --rm -p 8080:8080 -e APP_NAME=$(SVC) $(SVC):local


