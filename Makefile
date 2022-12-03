postgres:
	docker run --name postgres_local -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=admin -d postgres:12-alpine
createdb:
	docker exec -it postgres_local createdb --username=postgres --owner=postgres simple_bank
dropdb:
	docker exec -it postgres_local dropdb --username=postgres --owner=postgres simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://postgres:admin@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://postgres:admin@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
.PHONY: postgres createdb dropdb migrateup migratedown server
