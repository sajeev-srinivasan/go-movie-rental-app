package router

import (
	"github.com/gin-gonic/gin"
	"movie-rental-app/setup/config"
	"net/http"
)

func RegisterRoutes(engine *gin.Engine, config config.Config) {

	//dbConn := db.CreateConnection(config)
	movieGroup := engine.Group("/api/v1")
	{
		movieGroup.GET("/helloWorld", func(context *gin.Context) {
			context.JSON(http.StatusOK, "Hello, World!")
		})
	}
}
