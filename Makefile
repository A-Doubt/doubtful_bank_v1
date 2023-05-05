postgres:
	docker run --name postgres15 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root doubtful_bank
	docker exec -it postgres15 psql -U root -d doubtful_bank -c "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"

cleandb:
	docker exec -it postgres15 dropdb doubtful_bank
	docker exec -it postgres15 createdb --username=root --owner=root doubtful_bank
	docker exec -it postgres15 psql -U root -d doubtful_bank -c "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"
	migrate -path ./db/migration -database "postgresql://root:password@localhost:5433/doubtful_bank?sslmode=disable" -verbose up

dropdb:
	docker exec -it postgres15 dropdb doubtful_bank

sqlc:
	sqlc generate

migrateup:
	migrate -path ./db/migration -database "postgresql://root:password@localhost:5433/doubtful_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path ./db/migration -database "postgresql://root:password@localhost:5433/doubtful_bank?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test cleandb
