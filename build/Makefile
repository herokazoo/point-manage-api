# see https://postd.cc/auto-documented-makefile/
.PHONY: up down logs ps dry-migrate migrate help
.DEFAULT_GOAL := help

up: ## Do docker compose up with hot reload
	docker compose up -d

down: ## Do docker compose down
	docker compose down

logs: ## Tail docker compose logs
	docker compose logs -f

ps: ## Check container status
	docker compose ps

dry-migrate: ## Try migration using 'mysqldef'
	mysqldef -u point -p point -h 127.0.0.1 -P 33306 point --dry-run < ./db/schema.sql

migrate:  ## Execute migration using 'mysqldef'
	mysqldef -u point -p point -h 127.0.0.1 -P 33306 point < ./db/schema.sql

help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) |\
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
