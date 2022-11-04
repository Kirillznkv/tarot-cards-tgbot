include .env
export

.PHONY: build
build:
	go build -v ./cmd/main.go

.PHONY: test
test:
	go test -v -race ./...

.PHONY: db_up
db_up:
	docker run --rm --name my_postgres \
	-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
	-e POSTGRES_USER=$(DB_USER) \
	-e POSTGRES_DB=$(DB_NAME) \
	-p 5432:5432 -d postgres

.PHONY: db_down
db_down:
	docker stop my_postgres

.PHONY: migrate_up
migrate_up:
	goose -dir $(MIGRATIONS_DIR) postgres $(DB_URL) up

.PHONY: migrate_down
migrate_down:
	goose -dir $(MIGRATIONS_DIR) postgres $(DB_URL) down

.DEFAULT_GOAL := build