package database

import (
	"database/sql"
	"github.com/DavidHuie/gomigrate"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"nafue-api/config"
)

var (
	database *sql.DB
	migrator *gomigrate.Migrator
)

func init() {

	// create db connection
	db, err := sql.Open("mysql", config.DbUser+":"+config.DbPassword+"@"+config.DbServer+"/"+config.DbName+"?parseTime=true")
	if err != nil {
		log.Fatal("Database Error: ", err.Error())
	}
	database = db

	// create migrator
	m, _ := gomigrate.NewMigrator(database, gomigrate.Mysql{}, "database/migrations/")
	tableExists, _ := m.MigrationTableExists()
	if !tableExists {
		err = m.CreateMigrationsTable()
		if err != nil {
			log.Fatal("Migrator Error: ", err.Error())
		}
	}
	migrator = m

}

func Database() *sql.DB {
	return database
}

func Migrate() error {
	log.Println("Starting db migrations")

	err := migrator.Migrate()
	if err != nil {
		log.Println("Error migrating: ", err.Error())
		log.Println("Rolling back...")
		err = migrator.Rollback()
		return err
		if err != nil {
			log.Println("Error rolling back changes: ", err.Error())
			return err
		}
	}
	log.Println("Migrations finished.")
	return nil
}
