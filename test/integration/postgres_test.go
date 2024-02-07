package integration

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tc "github.com/testcontainers/testcontainers-go"
	"movie-rental-app/internal/app/model"
	"movie-rental-app/internal/app/router"
	"movie-rental-app/setup/config"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var db *sql.DB
var container tc.Container
var err error
var ctx context.Context

func TestMain(m *testing.M) {
	configs := config.InitConfigs("../../setup/env/local.yaml")
	migrationConfigs := configs.GetMigrationConfigs()
	container, db, err, ctx = createPostgresContainer(migrationConfigs.TestPath)
	if err != nil {
		return
	}
	code := m.Run()
	terminateContainer(container, ctx)
	os.Exit(code)
}

func TestShouldReturn200ResponseWithAllMovies(t *testing.T) {
	engine := gin.Default()
	router.RegisterRoutes(engine, db)

	request, err := http.NewRequest(http.MethodGet, "/api/v1/movies", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)

	var responseBody model.MovieResponse
	err = json.NewDecoder(responseRecorder.Body).Decode(&responseBody)
	fmt.Println("responseBody-->", responseBody)
	fmt.Println("err", err)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "success", responseBody.Status)
	assert.Equal(t, 5, len(responseBody.Data))
	assert.Equal(t, "tt0372784", responseBody.Data[0].Id)
}

func TestShouldReturn200ResponseWhenFilteringMoviesWithYear(t *testing.T) {
	engine := gin.Default()
	router.RegisterRoutes(engine, db)

	request, err := http.NewRequest(http.MethodGet, "/api/v1/movies?year=2010", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)

	var responseBody model.MovieResponse
	err = json.NewDecoder(responseRecorder.Body).Decode(&responseBody)
	fmt.Println("responseBody-->", responseBody)
	fmt.Println("err", err)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "success", responseBody.Status)
	assert.Equal(t, 2, len(responseBody.Data))
	assert.Equal(t, "tt1130884", responseBody.Data[0].Id)
	assert.Equal(t, "Shutter Island", responseBody.Data[0].Title)
	assert.Equal(t, "tt1375666", responseBody.Data[1].Id)
	assert.Equal(t, "Inception", responseBody.Data[1].Title)
}

func TestShouldReturn200ResponseWhenFetchingMovieById(t *testing.T) {
	engine := gin.Default()
	router.RegisterRoutes(engine, db)

	request, err := http.NewRequest(http.MethodGet, "/api/v1/movies/tt1130884", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)

	var responseBody model.Movie
	err = json.NewDecoder(responseRecorder.Body).Decode(&responseBody)
	fmt.Println("responseBody-->", responseBody)
	fmt.Println("err", err)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "tt1130884", responseBody.Id)
	assert.Equal(t, "Shutter Island", responseBody.Title)
	assert.Equal(t, 2010, responseBody.Year)
	assert.Equal(t, "Mystery", responseBody.Genre)
	assert.Equal(t, "Leonardo DiCaprio", responseBody.Actors)
}

func TestShouldReturn404WhenMovieIsNotAvailable(t *testing.T) {
	engine := gin.Default()
	router.RegisterRoutes(engine, db)

	request, err := http.NewRequest(http.MethodGet, "/api/v1/movies/abcd", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)

	var responseBody model.ErrorResponse
	err = json.NewDecoder(responseRecorder.Body).Decode(&responseBody)
	fmt.Println("responseBody-->", responseBody)
	fmt.Println("err", err)
	require.NoError(t, err)

	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
	assert.Equal(t, "no such movie is available", responseBody.Message)
}
