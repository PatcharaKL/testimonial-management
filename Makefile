# Makefile

DB_URL=postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable
MIGRATIONS_DIR=./migrations

.PHONY: migrate-up migrate-down migrate-force migrate-new

## Apply all up migrations
 migrate-up:
	bash -c 'set -a; source .env; set +a; DB_URL="postgres://$$DB_USER:$$DB_PASS@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable"; migrate -path $(MIGRATIONS_DIR) -database "$$DB_URL" up'

## Rollback last migration
 migrate-down:
	bash -c 'set -a; source .env; set +a; DB_URL="postgres://$$DB_USER:$$DB_PASS@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable"; migrate -path $(MIGRATIONS_DIR) -database "$$DB_URL" down 1'

## Force set migration version (careful: can break consistency)
 migrate-force:
	bash -c 'set -a; source .env; set +a; DB_URL="postgres://$$DB_USER:$$DB_PASS@$$DB_HOST:$$DB_PORT/$$DB_NAME?sslmode=disable"; migrate -path $(MIGRATIONS_DIR) -database "$$DB_URL" force $(VERSION)'

## Create a new migration file: make migrate-new name=create_users_table
 migrate-new:
	 bash -c 'set -a; source .env; set +a; migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)'
