sqlc:
	sqlc generate

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/gobank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/gobank?sslmode=disable" -verbose down

createdb:
	sudo docker exec -it postgres15alpl createdb --username=root --owner=root gobank

dropdb:
	sudo docker exec -it postgres15alpl dropdb gobank

migration:
	migrate create -ext sql -dir  db/migration -seq {theName}

test:
	go test -v -cover ./...

.PHONY: sqlc migrateup createdb dropdb test migration