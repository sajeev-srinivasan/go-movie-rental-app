package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHelloWorld(context *gin.Context) {
	context.JSON(http.StatusOK, "Hello, World!")
}
