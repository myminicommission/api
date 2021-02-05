all: clean test build

clean:
	go clean .

test:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

build:
	go build .

run:
	go run .

run-db:
	docker-compose up pgadmin

generate:
	go generate ./...