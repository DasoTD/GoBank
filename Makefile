sqlc:
	sqlc generate

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/test?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/test?sslmode=disable" -verbose down

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

.PHONY: sqlc migrateup createdb dropdb test migration server mockdb