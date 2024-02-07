package integration

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/golang-migrate/migrate/v4"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	db2 "movie-rental-app/internal/db"
)

var dbContainer tc.Container

func createPostgresContainer(path string) (tc.Container, *sql.DB, error, context.Context) {
	containerRequest := tc.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432"},
		Env: map[string]string{
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432"),
	}

	ctx := context.Background()
	dbContainer, err := tc.GenericContainer(ctx, tc.GenericContainerRequest{
		ContainerRequest: containerRequest,
		Started:          true,
	})
	if err != nil {
		println("Error creating postgres container: ", err.Error())
		return nil, nil, errors.New(fmt.Sprint("Error creating postgres container: ", err.Error())), ctx
	}

	host, err := dbContainer.Host(ctx)
	if err != nil {
		return nil, nil, errors.New(fmt.Sprint("Error fetching host: ", err.Error())), ctx
	}

	if err != nil {
		return nil, nil, err, nil
	}
	port, err := dbContainer.MappedPort(ctx, "5432")
	if err != nil {
		return nil, nil, errors.New(fmt.Sprint("Error fetching port: ", err.Error())), ctx
	}

	dataSourceName := makeDbUrl(host, port)
	fmt.Println("dataSourceName", dataSourceName)
	dbConn, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, nil, errors.New(fmt.Sprint("Error while connecting to database: ", err.Error())), ctx
	}

	migration, err := db2.GetMigration("test_movies", dbConn, path)
	if err != nil {
		return nil, nil, errors.New(fmt.Sprint("Error while creating migration instance: ", err.Error())), ctx
	}

	err = migration.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			println("No changes detected:", err.Error())
			return nil, nil, nil, ctx
		}
		println("Not able to run migration:", err.Error())
		return nil, nil, errors.New(fmt.Sprint("Not able to run migration: ", err.Error())), ctx
	}

	return dbContainer, dbConn, nil, ctx
}

func terminateContainer(container tc.Container, ctx context.Context) {
	err := container.Terminate(ctx)
	if err != nil {
		return
	}
}

func makeDbUrl(host string, port nat.Port) string {
	dbURI := fmt.Sprintf("postgres://postgres:postgres@%v:%v/testdb?sslmode=disable", host, port.Port())
	return dbURI
}
