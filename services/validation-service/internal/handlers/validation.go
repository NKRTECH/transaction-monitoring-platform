package handlers

import (
	"net/http"

	"github.com/gtrs/validation-service/internal/models"
	"github.com/gtrs/validation-service/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ValidationHandler handles validation-related endpoints
type ValidationHandler struct {
	validationService *services.ValidationService
}

// NewValidationHandler creates a new validation handler
func NewValidationHandler(validationService *services.ValidationService) *ValidationHandler {
	return &ValidationHandler{
		validationService: validationService,
	}
}

// ValidateTransaction handles transaction validation requests
func (h *ValidationHandler) ValidateTransaction(c *gin.Context) {
	var request models.ValidationRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		logrus.WithError(err).Error("Invalid validation request")
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"details": err.Error(),
		})
		return
	}

	// Validate the transaction
	result, err := h.validationService.ValidateTransaction(&request)
	if err != nil {
		logrus.WithError(err).Error("Validation failed")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Validation processing failed",
			"details": err.Error(),
		})
		return
	}

	logrus.WithFields(logrus.Fields{
		"transaction_id": request.TransactionID,
		"status":         result.Status,
	}).Info("Transaction validation completed")

	c.JSON(http.StatusOK, result)
}

// GetValidationResult retrieves a validation result by ID
func (h *ValidationHandler) GetValidationResult(c *gin.Context) {
	validationID := c.Param("id")

	if validationID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation ID is required",
		})
		return
	}

	result, err := h.validationService.GetValidationResult(validationID)
	if err != nil {
		logrus.WithError(err).WithField("validation_id", validationID).Error("Failed to retrieve validation result")
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Validation result not found",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}