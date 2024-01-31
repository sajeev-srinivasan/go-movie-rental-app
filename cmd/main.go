package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"movie-rental-app/internal/app/router"
	"movie-rental-app/internal/app/utils"
)

func main() {
	engine := gin.Default()
	router.RegisterRoutes(engine)

	var config utils.Config
	utils.GetConfig(&config)
	err := engine.Run(fmt.Sprint(config.Server.Host, ":", config.Server.Port))
	if err != nil {
		return
	}
	println("Listening and serving at 8080")
}
