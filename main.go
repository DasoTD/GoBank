package main

import (
	"database/sql"
	"log"

	"github.com/dasotd/gobank/api"
	db "github.com/dasotd/gobank/db/sqlc"
	"github.com/dasotd/gobank/util"
)




func main(){
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("can not load config")
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	bank := db.NewBank(conn)
	server :=api.NewServer(bank)
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start sever:", err)
	}

}

