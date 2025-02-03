run:
	docker-compose up --build

migrate:
	docker-compose exec app goose -dir migrations postgres "${DATABASE_URL}" up

