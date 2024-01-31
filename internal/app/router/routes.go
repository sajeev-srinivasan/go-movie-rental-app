package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(engine *gin.Engine) {
	movieGroup := engine.Group("/api/v1")
	{
		movieGroup.GET("/helloWorld", func(context *gin.Context) {
			context.JSON(http.StatusOK, "Hello, World!")
		})
	}
}
