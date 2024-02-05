package handlers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"movie-rental-app/internal/app/constants"
	"movie-rental-app/internal/app/model"
	"movie-rental-app/internal/app/service"
	"net/http"
)

type MovieHandler interface {
	GetMovies(ctx *gin.Context)
	GetMovie(ctx *gin.Context)
}

type movieHandler struct {
	movieService service.MovieService
}

func NewMovieHandler(movieService service.MovieService) MovieHandler {
	return &movieHandler{movieService: movieService}
}

func (m movieHandler) GetMovies(ctx *gin.Context) {
	year, _ := ctx.GetQuery("year")
	genre, _ := ctx.GetQuery("genre")
	actors, _ := ctx.GetQuery("actors")
	fmt.Println("---->", year, genre, actors)
	movies, err := m.movieService.GetMovies(year, genre, actors)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.MovieResponse{
			Response: model.Response{
				Status:  "error",
				Message: err.Error(),
			},
			Data: nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, model.MovieResponse{
		Response: model.Response{
			Status:  "success",
			Message: "",
		},
		Data: movies,
	})
}

func (m movieHandler) GetMovie(context *gin.Context) {
	movieId := context.Param("movieId")
	movie, err := m.movieService.GetMovie(movieId)
	if err != nil {
		if errors.Is(err, constants.ErrNoSuchMovie) {
			context.JSON(http.StatusNotFound, model.ErrorResponse{Message: err.Error()})
			return
		}
	}
	context.JSON(http.StatusOK, movie)
}
