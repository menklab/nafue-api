package database

import (
	"github.com/DavidHuie/gomigrate"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"github.com/menkveldj/nafue-api/config"
)

var (
	database *sqlx.DB
	migrator *gomigrate.Migrator
)

func init() {

	// create db connection
	connectionString := config.DbUser + ":" + config.DbPassword + "@" + config.DbServer + "/" + config.DbName + "?parseTime=true"
	db, err := sqlx.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("Database Error: ", err.Error())
	}
	database = db

	// create migrator
	m, _ := gomigrate.NewMigrator(database.DB, gomigrate.Mysql{}, "database/migrations/")
	tableExists, _ := m.MigrationTableExists()
	if !tableExists {
		err = m.CreateMigrationsTable()
		if err != nil {
			log.Fatal("Migrator Error: ", err.Error())
		}
	}
	migrator = m

}

func Database() *sqlx.DB {
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

func Rollback() error {
	log.Println("Starting db migrations")

	err := migrator.Rollback()
	if err != nil {
		log.Println("Error rollingback: ", err.Error())
		log.Println("Migrate up...")
		err = migrator.Migrate()
		return err
		if err != nil {
			log.Println("Error migrating up changes: ", err.Error())
			return err
		}
	}
	log.Println("Rollback finished.")
	return nil
}
