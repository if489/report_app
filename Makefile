.PHONY: db-init db-import build-and-run run all clean-up

db-init:
	@echo "Starting Postgres container"
	@docker run --name reports-postgres -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres
	@sleep 3
	@echo "Postgres succesfully launched."

db-import:
	@echo "Importing test data"
	@PGPASSWORD=password psql -h localhost -U postgres -d postgres -f create_reports.sql
	@echo "Test data imported succesfully"

build-and-run:
	go build -o reports_service main.go
	@chmod +x reports_service
	@echo "Server is starting at port 3000"
	@./reports_service

all: db-init db-import build-and-run

clean-up:
	docker rm -f reports-postgres

run:
	@echo "Server is starting at port 3000"
	go run main.go
