# Run Container docker
postgres:
	docker run --name db_postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

# Create database on container
createdb:
	docker exec -it db_postgres createdb --username=root --owner=root simple_bank

# Drop database
dropdb:
	docker exec -it db_postgres dropdb simple_bank

# Migrate up
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" --verbose up

#migrate down
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" --verbose down

# sqlc
sqlc:
	sqlc generate

# Running test
test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc