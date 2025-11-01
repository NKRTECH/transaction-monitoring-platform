package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gtrs/validation-service/internal/config"
	"github.com/gtrs/validation-service/internal/handlers"
	"github.com/gtrs/validation-service/internal/middleware"
	"github.com/gtrs/validation-service/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Setup logging
	setupLogging(cfg.LogLevel)

	// Initialize services
	validationService := services.NewValidationService()

	// Setup router
	router := setupRouter(cfg, validationService)

	// Create HTTP server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		logrus.WithFields(logrus.Fields{
			"port":    cfg.Port,
			"env":     cfg.Environment,
			"service": "validation-service",
		}).Info("Starting Validation Service")

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.WithError(err).Fatal("Failed to start server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Info("Shutting down server...")

	// Give outstanding requests 30 seconds to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logrus.WithError(err).Fatal("Server forced to shutdown")
	}

	logrus.Info("Server exited")
}

func setupLogging(level string) {
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	switch level {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}
}

func setupRouter(cfg *config.Config, validationService *services.ValidationService) *gin.Engine {
	// Set Gin mode based on environment
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// Add middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())
	router.Use(middleware.CORS())

	// Health endpoints
	healthHandler := handlers.NewHealthHandler()
	api := router.Group("/api")
	{
		health := api.Group("/health")
		{
			health.GET("", healthHandler.Health)
			health.GET("/ready", healthHandler.Ready)
			health.GET("/live", healthHandler.Live)
		}
	}

	// Validation endpoints (basic structure for now)
	validationHandler := handlers.NewValidationHandler(validationService)
	validation := api.Group("/validate")
	{
		validation.POST("", validationHandler.ValidateTransaction)
		validation.GET("/:id", validationHandler.GetValidationResult)
	}

	return router
}