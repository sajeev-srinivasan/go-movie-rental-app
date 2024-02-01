package db

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/lib/pq"
	"log"
	"movie-rental-app/setup/config"
)

func CreateConnection(config config.Config) *sql.DB {
	dataSourceName := fmt.Sprintf(
		"postgres://%s:%d/%s?user=%s&password=%s&sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
		config.Database.User,
		config.Database.Password,
	)

	dbConn, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal("unable to open connection with database ", err.Error())
	}
	if err := dbConn.Ping(); err != nil {
		log.Fatal("unable to ping database ", err.Error())
	}
	return dbConn
}

func NewMigration(config config.Config) (*migrate.Migrate, error) {
	dbConn := CreateConnection(config)

	driver, err := postgres.WithInstance(dbConn, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	migrateInstance, err := migrate.NewWithDatabaseInstance(
		"file://internal/db/migration",
		config.Database.Name,
		driver)
	if err != nil {
		return nil, err
	}
	return migrateInstance, nil
}
