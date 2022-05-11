.PHONY: postgres server build migration dropdb migrateup migratedown

postgres:
	docker compose up -d
server:
	go run main.go

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/lambda main.go

migration:
	migrate create -ext sql -dir db/migrations -seq init_schema 

dropdb:
	docker execc -it postgres12 dropdb bank_server

migrateup: 
	migrate -path db/migrations -database "postgres://admin:password@localhost:5432/default_db?sslmode=disable" up

migratedown:
	migrate -path db/migrations -database "postgres://admin:password@localhost:5432/default_db?sslmode=disable" down

dockerimage:
	# If facing an-error-failed-to-solve-with-frontend-dockerfile-v0
	# https://stackoverflow.com/a/66695181
	
	DOCKER_BUILDKIT=0 COMPOSE_DOCKER_CLI_BUILD=0 docker build . -t bank:1.0

