all: clean test build

clean:
	go clean .

test:
	go test ./...

build:
	go build .

run:
	go run .

generate:
	go generate ./...