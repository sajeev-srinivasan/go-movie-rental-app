package handlers

import (
	"encoding/json"
	"errors"
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

	movieRepository.On("GetAllMovies").Return([]model.Movie{
		{Id: "movie1", Title: "Harry Potter", Year: 2001, Genre: "Adventure", Actors: "Daniel Radcliff"},
		{Id: "movie2", Title: "Batman", Year: 2003, Genre: "Action", Actors: "Christian Bale"},
	}, nil)
	request, err := http.NewRequest(http.MethodGet, "/movies", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)

	var responseBody model.MovieResponse
	err = json.NewDecoder(responseRecorder.Body).Decode(&responseBody)
	fmt.Println("err", err)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "success", responseBody.Status)
	assert.Equal(t, 2, len(responseBody.Data))
	assert.Equal(t, "movie1", responseBody.Data[0].Id)
	assert.Equal(t, "movie2", responseBody.Data[1].Id)

	movieRepository.AssertNumberOfCalls(t, "GetAllMovies", 1)

}

func TestShouldReturn200ResponseWithEmptyListWhenFetchingAllMovies(t *testing.T) {
	engine := gin.Default()
	movieRepository := mocks.MovieRepository{}
	movieService := service.NewMovieService(&movieRepository)
	movieHandler := NewMovieHandler(movieService)
	engine.GET("/movies", movieHandler.GetMovies)

	movieRepository.On("GetAllMovies").Return([]model.Movie{}, nil)
	request, err := http.NewRequest(http.MethodGet, "/movies", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)

	var responseBody model.MovieResponse
	err = json.NewDecoder(responseRecorder.Body).Decode(&responseBody)
	fmt.Println("err", err)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "success", responseBody.Status)
	assert.Equal(t, 0, len(responseBody.Data))
	assert.Equal(t, []model.Movie{}, responseBody.Data)

	movieRepository.AssertNumberOfCalls(t, "GetAllMovies", 1)
}

func TestShouldReturn500ResponseWhenInternalServerErrorOccurs(t *testing.T) {
	engine := gin.Default()

	movieRepository := mocks.MovieRepository{}
	movieService := service.NewMovieService(&movieRepository)
	movieHandler := NewMovieHandler(movieService)

	engine.GET("/movies", movieHandler.GetMovies)

	movieRepository.On("GetAllMovies").Return([]model.Movie{}, errors.New("unable to connect to database"))
	request, err := http.NewRequest(http.MethodGet, "/movies", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)

	var responseBody model.MovieResponse
	err = json.NewDecoder(responseRecorder.Body).Decode(&responseBody)
	fmt.Println("err", err)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, responseRecorder.Code)
	assert.Equal(t, "error", responseBody.Status)
	assert.Equal(t, "unable to connect to database", responseBody.Message)

	movieRepository.AssertNumberOfCalls(t, "GetAllMovies", 1)
}

func TestShouldReturnMoviesWhenFilteringWithYear(t *testing.T) {
	engine := gin.Default()

	movieRepository := mocks.MovieRepository{}
	movieService := service.NewMovieService(&movieRepository)
	movieHandler := NewMovieHandler(movieService)

	engine.GET("/movies", movieHandler.GetMovies)

	movieRepository.On("GetMovies", "2003", "", "").Return([]model.Movie{
		{Id: "movie2", Title: "Batman", Year: 2003, Genre: "Action", Actors: "Christian Bale"},
	}, nil)
	request, err := http.NewRequest(http.MethodGet, "/movies?year=2003", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)

	var responseBody model.MovieResponse
	fmt.Println(responseRecorder.Body)
	err = json.NewDecoder(responseRecorder.Body).Decode(&responseBody)
	fmt.Println("err", err)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "success", responseBody.Status)
	assert.Equal(t, 1, len(responseBody.Data))
	assert.Equal(t, "movie2", responseBody.Data[0].Id)
}
