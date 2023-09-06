db-migrate:
	docker-compose run --rm db-migrate sh -c 'migrate -source migrations -database $$MYSQL_URL up'

teardown:
	docker-compose down -v

init: mysql sleep db-migrate

mysql:
	docker-compose up -d mysql

sleep:
	sleep 5