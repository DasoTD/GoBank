sqlc:
	sqlc generate

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/test?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/test?sslmode=disable" -verbose down

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/test?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/test?sslmode=disable" -verbose down 1


createdb:
	sudo docker exec -it postgres15alpl createdb --username=root --owner=root test

dropdb:
	sudo docker exec -it postgres15alpl dropdb test

migration:
	migrate create -ext sql -dir  db/migration -seq {theName}

test:
	go test -v -cover ./...
server:
	go run main.go

mockdb:
	mockgen --build_flags=--mod=mod -package mockdb -destination db/mock/bank.go github.com/dasotd/gobank/db/sqlc Bank

.PHONY: sqlc migrateup migratedown migrateup1 migratedown1 createdb dropdb test migration server mockdb