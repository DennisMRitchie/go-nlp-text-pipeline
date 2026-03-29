.PHONY: run build docker-up test

run:
	go run ./cmd/api

build:
	go build -o bin/nlp-pipeline ./cmd/api

docker-up:
	docker-compose up --build

test:
	go test ./... -v
