package handlers

import (
	"github.com/gin-gonic/gin"
	"movie-rental-app/internal/app/model"
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
	movies, err := m.movieService.GetMovies()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.MovieResponse{
			Status:  "error",
			Message: err.Error(),
			Data:    nil,
		})
	}
	ctx.JSON(http.StatusOK, model.MovieResponse{
		Status:  "success",
		Message: "",
		Data:    movies,
	})
}
