install-tools:
	go install github.com/cosmtrek/air@latest

migrate:
	docker-compose run --rm migrate -path /migrations -database "mysql://user:password@tcp(db:3306)/go_template" up

# Syntax: make migrate-create MIGRATE_NAME=...
migrate-create:
	docker-compose run --rm migrate create -ext sql -dir /migrations $(MIGRATE_NAME)

gen-mocks:
	docker-compose run -w /src mockery --all

teardown:
	docker-compose down -v

init: db gen-mocks sleep migrate

db:
	docker-compose up -d db

sleep:
	sleep 10

dev:
	air

test:
	go test -cover ./...

.PHONY: install-tools migrate migrate-create gen-mocks teardown init db dev sleep test