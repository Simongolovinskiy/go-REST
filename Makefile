DATABASE_URL=postgres://user:password@go_db:5432/finance_db?sslmode=disable

run:
	docker-compose up --build

migrate:
	docker-compose exec app goose -dir migrations postgres ${DATABASE_URL} up

