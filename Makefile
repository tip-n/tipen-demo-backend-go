#!make
include .env

db_url := postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}

docker-up:
	docker compose up --build

docker-down:
	docker compose down --volumes

run:
	go run cmd/main.go

seed-restaurant:
	go run cmd/seeds/restaurant/main.go

build:
	cd cmd/server && go build -tags netgo -ldflags '-s -w' -o ../../app

start:
	./app

generate-migration:
	@read -p "Migration name : " name; \
	migrate create -ext sql -dir db/migrations -seq $$name

force-migrate:
	@read -p "Migration ver : " version; \
	migrate -database "${db_url}?sslmode=disable" -path db/migrations force $$version


migrate:
	migrate -database "${db_url}?sslmode=disable" -path db/migrations down
	migrate -database "${db_url}?sslmode=disable" -path db/migrations up

migrate-down:
	migrate -database "${db_url}?sslmode=disable" -path db/migrations down

generate:
	mkdir generated || true
	oapi-codegen -package schema schema.yml > generated/schema.gen.go