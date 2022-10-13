.PHONY: postgres server build migration dropdb migrateup migratedown createdb

postgres:
	docker compose up -d
server:
	go run main.go

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/lambda main.go

migration:
	migrate create -ext sql -dir models/migrations -seq init_schema

dropdb:
	docker exec -it postgres14 dropdb --username=admin bank_server

migrateup: 
	migrate -path models/migrations -database "postgres://admin:password@localhost:5432/bank_server?sslmode=disable" up

migratedown:
	migrate -path models/migrations -database "postgres://admin:password@localhost:5432/bank_server?sslmode=disable" down

dockerimage:
	# If facing an-error-failed-to-solve-with-frontend-dockerfile-v0
	# https://stackoverflow.com/a/66695181
	
	DOCKER_BUILDKIT=0 COMPOSE_DOCKER_CLI_BUILD=0 docker build . -t grofffer/bank:1.1

createdb:
	docker exec -it postgres14 createdb --username=admin --owner=admin bank_server

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

backupdb:
	docker exec (container_name) pg_dump -U (db_user) (db_schema) > backup.sql

mockgen:
	# go install github.com/golang/mock/mockgen@v1.6.0
	mockgen -source models/store/store.go -package store -destination=models/store/mock_store.go