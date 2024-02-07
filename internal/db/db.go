package db

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
	"movie-rental-app/setup/config"
)

func CreateConnection(config config.DatabaseConfigs) *sql.DB {
	dataSourceName := fmt.Sprintf(
		"postgres://%s:%d/%s?user=%s&password=%s&sslmode=disable",
		config.Host,
		config.Port,
		config.DatabaseName,
		config.User,
		config.Password,
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

func NewMigration(dbConfig config.DatabaseConfigs, migrationConfigs config.MigrationConfigs) (*migrate.Migrate, error) {
	dbConn := CreateConnection(dbConfig)
	return GetMigration(dbConfig.DatabaseName, dbConn, migrationConfigs.DevPath)
}

func GetMigration(dbName string, dbConn *sql.DB, path string) (*migrate.Migrate, error) {
	fmt.Println("after create conn")
	driver, err := postgres.WithInstance(dbConn, &postgres.Config{})

	if err != nil {
		fmt.Println("WithInstance ->", err)
		return nil, err
	}
	fmt.Println("after with instance")
	migrateInstance, err := migrate.NewWithDatabaseInstance(
		"file://"+path,
		dbName,
		driver)
	if err != nil {
		fmt.Println("NewWithDatabaseInstance ->", err)
		return nil, err
	}
	fmt.Println("after NewWithDatabaseInstance  line 51")
	return migrateInstance, nil
}
