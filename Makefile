TEST_DIR = test
MIGRATIONS_DIR = migrations
CONTAINER_NAME = pg_migration_test
DB_URL = "postgres://postgres:postgres@127.0.0.1:5432/bbj_db?sslmode=disable"

.PHONY: run stop clean status logs migrate

run:
	cd $(TEST_DIR) && podman compose up -d
	podman ps --filter name=$(CONTAINER_NAME)
	podman logs -f $(CONTAINER_NAME)

stop:
	cd $(TEST_DIR) && podman compose down

clean:
	cd $(TEST_DIR) && podman compose down
	sudo rm -rf $(TEST_DIR)/_data

status:
	podman ps --filter name=$(CONTAINER_NAME)

logs:
	podman logs -f $(CONTAINER_NAME)

migrate:
	goose -dir $(MIGRATIONS_DIR) postgres $(DB_URL) up
