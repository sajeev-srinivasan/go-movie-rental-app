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
var configs config.Config

func TestMain(m *testing.M) {
	config.GetConfig(&configs, "../../setup/env/inttest.yaml")
	container, db, err, ctx = createPostgresContainer(configs)
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
	assert.Equal(t, 0, len(responseBody.Data))
}
