package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"recommendation/api"
	"github.com/stretchr/testify/assert"
)

func TestGetOrigamiOfTheDay(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Setup the router
	router := gin.Default()
	router.GET("/api/origami-of-the-day", api.GetOrigamiOfTheDay)

	// Create a response recorder
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/origami-of-the-day", nil)
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
}
