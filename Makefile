install-tools:
	go install github.com/cosmtrek/air@latest

migrate:
	docker-compose run --rm migrate -path /migrations -database "mysql://user:password@tcp(db:3306)/go_template" up

# Syntax: make migrate-create MIGRATE_NAME=...
migrate-create:
	docker-compose run --rm migrate create -ext sql -dir /migrations $(MIGRATE_NAME)

teardown:
	docker-compose down -v

init: db sleep migrate

db:
	docker-compose up -d db

dev:
	air

sleep:
	sleep 10

.PHONY: install-tools migrate migrate-create teardown init db dev sleep