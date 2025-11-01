package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthHandler handles health check endpoints
type HealthHandler struct{}

// NewHealthHandler creates a new health handler
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthResponse represents the health check response
type HealthResponse struct {
	Status    string    `json:"status"`
	Service   string    `json:"service"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version,omitempty"`
}

// Health handles the main health check endpoint
func (h *HealthHandler) Health(c *gin.Context) {
	response := HealthResponse{
		Status:    "UP",
		Service:   "validation-service",
		Timestamp: time.Now(),
		Version:   "1.0.0-SNAPSHOT",
	}

	c.JSON(http.StatusOK, response)
}

// Ready handles the readiness probe endpoint
func (h *HealthHandler) Ready(c *gin.Context) {
	// TODO: Add actual readiness checks (database connectivity, etc.)
	response := HealthResponse{
		Status:    "READY",
		Service:   "validation-service",
		Timestamp: time.Now(),
	}

	c.JSON(http.StatusOK, response)
}

// Live handles the liveness probe endpoint
func (h *HealthHandler) Live(c *gin.Context) {
	response := HealthResponse{
		Status:    "ALIVE",
		Service:   "validation-service",
		Timestamp: time.Now(),
	}

	c.JSON(http.StatusOK, response)
}