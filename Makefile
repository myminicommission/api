all: bin/api
test: unit-test

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

generate:
	go generate ./...