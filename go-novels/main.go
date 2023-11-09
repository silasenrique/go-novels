package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"go-models/api"
)

func main() {
	db, err := sql.Open("sqlite3", "/home/senrique/novels.db")
	if err != nil {
		log.Fatalf("nao foi possivel se conectar no banco. err: %s", err)
	}

	api.NewServer(db).
		SetRoutes().
		Run(":8080")
}
