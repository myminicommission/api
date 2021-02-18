all: lint test bin/api
test: unit-test
start: lint test bin/api start/bin

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

.PHONY: start/db
start/db:
	docker-compose up pgadmin postgres

.PHONY: stop/db
stop/db:
	docker-compose stop pgadmin postgres