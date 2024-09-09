.PHONY: generate help

generate:
	sqlc generate

help:
	@echo "Available commands:"
	@echo "  make generate  - Generate SQL code using sqlc"
	@echo "  make help      - Show this help message"
	@echo "  make up        - Start the containers"
	@echo "  make down      - Stop the containers"

up:
	docker compose up -d

down:
	docker compose down
