#!make
include .env

db_url := postgres://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}

docker-up:
	sudo docker compose up -d

docker-down:
	sudo docker compose down

run:
	go run cmd/server/run.go

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
