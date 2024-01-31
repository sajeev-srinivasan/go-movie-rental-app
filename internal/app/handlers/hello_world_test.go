package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorldShouldReturn200(t *testing.T) {
	engine := gin.Default()

	engine.GET("/helloWorld", GetHelloWorld)

	request, err := http.NewRequest(http.MethodGet, "/helloWorld", nil)
	require.NoError(t, err)

	responseRecorder := httptest.NewRecorder()
	engine.ServeHTTP(responseRecorder, request)

	var responseBody string
	err = json.NewDecoder(responseRecorder.Body).Decode(&responseBody)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "Hello, World!", responseBody)
}
