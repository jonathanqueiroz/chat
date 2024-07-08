.PHONY: build run test clean up stop

build:
	go build -o main ./cmd/server/main.go

run: build
	./main

test:
	go test -v ./...

clean:
	go clean
	rm -f main

docker-build:
	docker compose build

docker-run:
	docker compose up -d

up: docker-run

stop:
	docker compose down

db-up:
	docker compose up -d postgres
