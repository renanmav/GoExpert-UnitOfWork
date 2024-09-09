.PHONY: generate help

generate:
	sqlc generate

help:
	@echo "Available commands:"
	@echo "  make generate  - Generate SQL code using sqlc"
	@echo "  make help      - Show this help message"