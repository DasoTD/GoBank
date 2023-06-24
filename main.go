package main

import (
	"database/sql"
	"log"

	"github.com/dasotd/gobank/api"
	db "github.com/dasotd/gobank/db/sqlc"
)

const (
	DBDriver ="postgres"
	DBSource = "postgresql://root:secret@localhost:5432/test?sslmode=disable"
	serverAddress= "0.0.0.0:8080"
)


func main(){
	conn, err := sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	bank := db.NewBank(conn)
	server :=api.NewServer(bank)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start sever:", err)
	}

}

