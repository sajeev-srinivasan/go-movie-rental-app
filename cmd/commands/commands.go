package commands

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"movie-rental-app/internal/app/router"
	"movie-rental-app/internal/db"
	"movie-rental-app/setup/config"
)

const (
	initHttpServeCommand  = "http-serve"
	initMigrationsCommand = "migrate"
)

func Execute(cmd, configFile string) {
	run, ok := commands()[cmd]
	if !ok {
		return
	}
	run(configFile)
}

func commands() map[string]func(configFile string) {
	return map[string]func(configFile string){
		initHttpServeCommand:  initHttpServer,
		initMigrationsCommand: initMigrations,
	}
}

func initMigrations(file string) {
	configs := config.InitConfigs(file)
	databaseConfigs := configs.GetDatabaseConfigs()
	migrationConfigs := configs.GetMigrationConfigs()
	migrateInstance, err := db.NewMigration(databaseConfigs, migrationConfigs)
	if err != nil {
		println("error in fetching migration instance:", err)
		return
	}
	err = migrateInstance.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			println("No changes detected:", err.Error())
			return
		}
		println("Not able to run migration:", err.Error())
		return
	}
}

func initHttpServer(file string) {
	configs := config.InitConfigs(file)
	serverConfigs := configs.GetServerConfigs()
	databaseConfigs := configs.GetDatabaseConfigs()
	engine := gin.Default()
	dbConn := db.CreateConnection(databaseConfigs)
	router.RegisterRoutes(engine, dbConn)
	err := engine.Run(fmt.Sprint(serverConfigs.Host, ":", serverConfigs.Port))
	if err != nil {
		return
	}

	println("Listening and serving to 8080...")
}
