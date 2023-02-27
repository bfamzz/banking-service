postgres:
	docker run --name banking-service-db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it banking-service-db createdb --username=root --owner=root banking-service

dropdb:
	docker exec -it banking-service-db dropdb --username=root --owner=root banking-service

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/banking-service?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/banking-service?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server