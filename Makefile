start: start/api
down:
	docker-compose down

.PHONY: generate
generate:
	go generate ./...

.PHONY: start/api
start/api:
	docker-compose --compatibility up --build api

.PHONY: start/db
start/db:
	docker-compose up postgres

.PHONY: stop/db
stop/db:
	docker-compose stop postgres

.PHONY: bin/api
bin/api:
	@docker build --target build .