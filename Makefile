all: lint test bin/api
test: unit-test
start: lint test start/api
down:
	docker-compose down

PLATFORM=local

.PHONY: bin/api
bin/api:
	@docker build --target bin --output bin/ --platform ${PLATFORM} .

.PHONY: unit-test
unit-test:
	@docker build . --target unit-test

.PHONY: lint
lint:
	@docker build . --target lint

.PHONY: generate
generate:
	go generate ./...

.PHONY: start/bin
start/bin:
	bin/api

.PHONY: start/api
start/api:
	docker-compose --compatibility up --build api

.PHONY: start/db
start/db:
	docker-compose up postgres

.PHONY: stop/db
stop/db:
	docker-compose stop postgres