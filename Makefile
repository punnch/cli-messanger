include .env
export 

export PROJECT_ROOT=$(shell pwd)

# env
env-up:
	@docker compose up -d postgres

env-down:
	@docker compose down postgres
	
env-cleanup:
	@read -p "Do you want to clean up your environment files? You may lose your data. [y/N]: " ans; \
	if [ "$$ans" == "y" ]; then \
		docker compose down postgres && \
		sudo rm -rf ~/docker-volumes/messanger-pgdata && \
		echo "Environment files cleaned up successfuly"; \
	else \
		echo "Cleaning of environment files canceled"; \
	fi

env-port-forward:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder

logs-cleanup:
	@read -p "Do you want to clean up your logs? You may lose your data. [y/N]: " ans; \
	if [ "$$ans" == "y" ]; then \
		sudo rm -rf ${PROJECT_ROOT}/out/logs && \
		echo "Logs cleaned up successfuly"; \
	else \
		echo "Cleaning of logs canceled"; \
	fi

# migrate
migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "seq not set. Example: make migrate-create seq=init"; \
		exit 1; \
	fi; 
	docker compose run --rm --user $(shell id -u):$(shell id -g) postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "action not set. Example: make migrate-action action=up"; \
		exit 1; \
	fi; \
	docker compose run --rm postgres-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@postgres:5432/${POSTGRES_DB}?sslmode=disable \
		$(action)

# app
messanger-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/server/main.go

messanger-deploy:
	@docker compose up -d --build messanger

messanger-undeploy:
	@docker compose down messanger

# client-side
client-run:
	@go run ${PROJECT_ROOT}/cmd/client/main.go

# other
ps:
	@docker compose ps
