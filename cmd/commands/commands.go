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
	var configs config.Config
	config.GetConfig(&configs, file)
	migrateInstance, err := db.NewMigration(configs)
	if err != nil {
		println("error in fetching migration instance:", err)
		return
	}
	err = migrateInstance.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			println("No changes detected:", err)
			return
		}
		println("Not able to run migration:", err)
		return
	}
}

func initHttpServer(file string) {
	engine := gin.Default()
	var configs config.Config
	config.GetConfig(&configs, file)
	router.RegisterRoutes(engine, configs)
	err := engine.Run(fmt.Sprint(configs.Server.Host, ":", configs.Server.Port))
	if err != nil {
		return
	}

	println("Listening and serving to 8080...")
}
