package config

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	dsn := "root:root@tcp(localhost:3306)/url_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
