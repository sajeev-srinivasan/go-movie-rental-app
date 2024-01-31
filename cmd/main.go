package main

import (
	"github.com/gin-gonic/gin"
	"movie-rental-app/internal/app/router"
)

func main() {
	engine := gin.Default()
	router.RegisterRoutes(engine)
	err := engine.Run("localhost:8080")
	if err != nil {
		return
	}
	println("Listening and serving at 8080")
}
