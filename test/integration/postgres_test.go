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
	"movie-rental-app/internal/app/handlers"
	"movie-rental-app/internal/app/model"
	"movie-rental-app/internal/app/repository"
	"movie-rental-app/internal/app/service"
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
	container, db, err, ctx = createPostgresContainer()
	if err != nil {
		return
	}
	code := m.Run()
	terminateContainer(container, ctx)
	os.Exit(code)
}

func TestShouldReturn200ResponseWithAllMovies(t *testing.T) {
	engine := gin.Default()

	movieRepository := repository.NewMovieRepository(db)
	movieService := service.NewMovieService(movieRepository)
	movieHandler := handlers.NewMovieHandler(movieService)
	engine.GET("/movies", movieHandler.GetMovies)

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
}
