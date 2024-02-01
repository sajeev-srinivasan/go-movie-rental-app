package handlers

import (
	"github.com/gin-gonic/gin"
	"movie-rental-app/internal/app/service"
	"net/http"
)

type MovieHandler interface {
	GetMovies(ctx *gin.Context)
}

type movieHandler struct {
	movieService service.MovieService
}

func NewMovieHandler(movieService service.MovieService) MovieHandler {
	return &movieHandler{movieService: movieService}
}

func (m movieHandler) GetMovies(ctx *gin.Context) {
	movies, _ := m.movieService.GetMovies()
	ctx.JSON(http.StatusOK, movies)
}
