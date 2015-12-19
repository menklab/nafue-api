package database

import ("database/sql"
		_ "github.com/lib/pq"
 		"log"
)
var (
	database			*sql.DB
	databaseConfig		string
)

func init() {
	db, err := sql.Open("postgres", "user=sparticus dbname=sparticus password=sparticus sslmode=require")
	if err != nil {
		log.Fatal("Database Error: ", err.Error())
	}
	database = db
}

func Database() *sql.DB {
	return database
}