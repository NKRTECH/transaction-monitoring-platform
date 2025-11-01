package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthHandler_Health(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create handler
	handler := NewHealthHandler()

	// Create test router
	router := gin.New()
	router.GET("/api/health", handler.Health)

	// Create test request
	req, _ := http.NewRequest("GET", "/api/health", nil)
	w := httptest.NewRecorder()

	// Perform request
	router.ServeHTTP(w, req)

	// Assert response
	assert.Equal(t, http.StatusOK, w.Code)

	var response HealthResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "UP", response.Status)
	assert.Equal(t, "validation-service", response.Service)
	assert.Equal(t, "1.0.0-SNAPSHOT", response.Version)
}

func TestHealthHandler_Ready(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create handler
	handler := NewHealthHandler()

	// Create test router
	router := gin.New()
	router.GET("/api/health/ready", handler.Ready)

	// Create test request
	req, _ := http.NewRequest("GET", "/api/health/ready", nil)
	w := httptest.NewRecorder()

	// Perform request
	router.ServeHTTP(w, req)

	// Assert response
	assert.Equal(t, http.StatusOK, w.Code)

	var response HealthResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "READY", response.Status)
	assert.Equal(t, "validation-service", response.Service)
}

func TestHealthHandler_Live(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create handler
	handler := NewHealthHandler()

	// Create test router
	router := gin.New()
	router.GET("/api/health/live", handler.Live)

	// Create test request
	req, _ := http.NewRequest("GET", "/api/health/live", nil)
	w := httptest.NewRecorder()

	// Perform request
	router.ServeHTTP(w, req)

	// Assert response
	assert.Equal(t, http.StatusOK, w.Code)

	var response HealthResponse
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "ALIVE", response.Status)
	assert.Equal(t, "validation-service", response.Service)
}