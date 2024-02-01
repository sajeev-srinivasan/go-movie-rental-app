package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"movie-rental-app/internal/app/model"
	"movie-rental-app/internal/app/repository/mocks"
	"movie-rental-app/internal/app/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShouldReturn200ResponseWithMovieListWhenFetchingAllMovies(t *testing.T) {
	engine := gin.Default()
	movieRepository := mocks.MovieRepository{}
	movieService := service.NewMovieService(&movieRepository)
	movieHandler := NewMovieHandler(movieService)
	engine.GET("/movies", movieHandler.GetMovies)

	movieRepository.On("GetMovies").Return([]model.Movie{
		{Id: "movie1", Title: "Harry Potter", Year: 2001, Genre: "Adventure", Actors: "Daniel Radcliff"},
		{Id: "movie2", Title: "Batman", Year: 2003, Genre: "Action", Actors: "Christian Bale"},
	}, nil)
	request, err := http.NewRequest(http.MethodGet, "/movies", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)

	var responseBody []model.Movie
	err = json.NewDecoder(responseRecorder.Body).Decode(&responseBody)
	fmt.Println("err", err)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, 2, len(responseBody))
	assert.Equal(t, "movie1", responseBody[0].Id)
	assert.Equal(t, "movie2", responseBody[1].Id)

	movieRepository.AssertNumberOfCalls(t, "GetMovies", 1)

}
func TestShouldReturn200ResponseWithEmptyListWhenFetchingAllMovies(t *testing.T) {
	engine := gin.Default()
	movieRepository := mocks.MovieRepository{}
	movieService := service.NewMovieService(&movieRepository)
	movieHandler := NewMovieHandler(movieService)
	engine.GET("/movies", movieHandler.GetMovies)

	movieRepository.On("GetMovies").Return([]model.Movie{}, nil)
	request, err := http.NewRequest(http.MethodGet, "/movies", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)

	var responseBody []model.Movie
	err = json.NewDecoder(responseRecorder.Body).Decode(&responseBody)
	fmt.Println("err", err)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, 0, len(responseBody))
	assert.Equal(t, []model.Movie{}, responseBody)

	movieRepository.AssertNumberOfCalls(t, "GetMovies", 1)

}
