start: start/api
down:
	@docker compose down

.PHONY: generate
generate:
	go get github.com/99designs/gqlgen/cmd
	go generate ./...

.PHONY: start/api
start/api:
	@docker compose up --build api

.PHONY: start/db
start/db:
	@docker compose up postgres

.PHONY: stop/db
stop/db:
	@docker compose stop postgres

.PHONY: test-lint-and-build
test-lint-and-build:
	@docker build --target build .

.PHONY: clean
clean:
	rm -rf graph/generated
	rm -rf graph/model
	rm -rf bin
	rm -f coverage.out
