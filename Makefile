SVC=yourservice
BIN=bin/$(SVC)

.PHONY: build run tidy test lint fmt docker docker-run

build:
	@mkdir -p bin
	GO111MODULE=on CGO_ENABLED=0 go build -o $(BIN) ./cmd/service

run: build
	APP_NAME=$(SVC) $(BIN)

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


