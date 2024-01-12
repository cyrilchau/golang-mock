.PHONY: help database-up database-down migration-up migration-down local run

help:
	@echo "Available targets:"
	@echo "  make database-up   	- Start the database container"
	@echo "  make database-down 	- Stop and remove the database container"
	@echo "  make migration-up  	- Run database migrations"
	@echo "  make migration-down 	- Rollback database migrations"
	@echo "  make local         	- Run the application locally"
	@echo "  make run           	- Start the database, run migrations, and start the application locally"
	@echo "  make down           	- Shutdown the database and down migrations"



# Directory where migration files are located
MIGRATION_DIR := database/postgres/migration

# This target waits for the Postgres container to become available
wait-for-postgres:
	@echo "Waiting for Postgres container to start..."
	@until docker compose exec pg-db psql "postgresql://admin:PAssw0rd@127.0.0.1:5432/postgres" -c '\q'; do \
		sleep 3; \
	done
	@echo "Postgres is up and running!"
	
database-up: 
	docker compose up pg-db -d

# service-up:
# 	docker compose up golang-clean-architecture -d

docker-down:
	docker compose down 

migration-up: wait-for-postgres
	GOOSE_DRIVER=postgres GOOSE_DBSTRING="user=admin dbname=postgres sslmode=disable" goose -dir=$(MIGRATION_DIR) up

migration-down: 
	GOOSE_DRIVER=postgres GOOSE_DBSTRING="user=admin dbname=postgres sslmode=disable" goose -dir=$(MIGRATION_DIR) down

run: database-up
# run: database-up migration-up service-up

down: docker-down
# down : migration-down docker-down

# mock-repository:
# 	mockgen -source internal/users/repository/repository.go -destination internal/users/mock/repository_mock.go -package=mocks

# mock-usecase:
# 	mockgen -source internal/users/usecase/usecase.go -destination internal/users/mock/usecase_mock.go -package=mocks

