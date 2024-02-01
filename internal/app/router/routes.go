package router

import (
	"github.com/gin-gonic/gin"
	"movie-rental-app/internal/app/handlers"
	"movie-rental-app/internal/app/repository"
	"movie-rental-app/internal/app/service"
	"movie-rental-app/internal/db"
	"movie-rental-app/setup/config"
	"net/http"
)

func RegisterRoutes(engine *gin.Engine, config config.Config) {

	dbConn := db.CreateConnection(config)
	movieRepository := repository.NewMovieRepository(dbConn)
	movieService := service.NewMovieService(movieRepository)
	movieHandler := handlers.NewMovieHandler(movieService)
	movieGroup := engine.Group("/api/v1")
	{
		movieGroup.GET("/helloWorld", func(context *gin.Context) {
			context.JSON(http.StatusOK, "Hello, World!")
		})
		movieGroup.GET("/movies", movieHandler.GetMovies)
	}
}
