package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	database       *sql.DB
	databaseConfig string
)

func init() {
	db, err := sql.Open("mysql", "sparticus:password@/sparticus")
	if err != nil {
		log.Fatal("Database Error: ", err.Error())
	}
	database = db
}

func Database() *sql.DB {
	return database
}
