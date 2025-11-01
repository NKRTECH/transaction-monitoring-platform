package services

import (
	"testing"
	"time"

	"github.com/gtrs/validation-service/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestValidationService_ValidateTransaction_Success(t *testing.T) {
	service := NewValidationService()

	request := &models.ValidationRequest{
		TransactionID: "test-txn-123",
		Type:          "PAYMENT",
		Amount:        1000.00,
		Currency:      "USD",
		Counterparty: models.Counterparty{
			ID:   "cp-123",
			Name: "Test Corp",
			Type: "BUSINESS",
		},
		Timestamp: time.Now(),
	}

	result, err := service.ValidateTransaction(request)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, request.TransactionID, result.TransactionID)
	assert.Equal(t, models.ValidationStatusPassed, result.Status)
	assert.NotEmpty(t, result.ID)
	assert.True(t, len(result.Rules) > 0)
	assert.True(t, result.ProcessingTime >= 0)
}

func TestValidationService_ValidateTransaction_AmountExceedsLimit(t *testing.T) {
	service := NewValidationService()

	request := &models.ValidationRequest{
		TransactionID: "test-txn-456",
		Type:          "PAYMENT",
		Amount:        2000000.00, // Exceeds default limit of 1,000,000
		Currency:      "USD",
		Counterparty: models.Counterparty{
			ID:   "cp-456",
			Name: "Test Corp",
			Type: "BUSINESS",
		},
		Timestamp: time.Now(),
	}

	result, err := service.ValidateTransaction(request)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, models.ValidationStatusFailed, result.Status)
	assert.Equal(t, "VALIDATION_FAILED", result.ErrorCode)
	assert.NotEmpty(t, result.ErrorMessage)

	// Check that amount limit rule failed
	amountRuleFailed := false
	for _, rule := range result.Rules {
		if rule.RuleID == "amount-limit" && rule.Status == "FAILED" {
			amountRuleFailed = true
			break
		}
	}
	assert.True(t, amountRuleFailed, "Amount limit rule should have failed")
}

func TestValidationService_ValidateTransaction_InvalidCurrency(t *testing.T) {
	service := NewValidationService()

	request := &models.ValidationRequest{
		TransactionID: "test-txn-789",
		Type:          "PAYMENT",
		Amount:        1000.00,
		Currency:      "XYZ", // Invalid currency
		Counterparty: models.Counterparty{
			ID:   "cp-789",
			Name: "Test Corp",
			Type: "BUSINESS",
		},
		Timestamp: time.Now(),
	}

	result, err := service.ValidateTransaction(request)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, models.ValidationStatusFailed, result.Status)

	// Check that currency rule failed
	currencyRuleFailed := false
	for _, rule := range result.Rules {
		if rule.RuleID == "currency-check" && rule.Status == "FAILED" {
			currencyRuleFailed = true
			break
		}
	}
	assert.True(t, currencyRuleFailed, "Currency rule should have failed")
}

func TestValidationService_ValidateTransaction_MissingCounterparty(t *testing.T) {
	service := NewValidationService()

	request := &models.ValidationRequest{
		TransactionID: "test-txn-000",
		Type:          "PAYMENT",
		Amount:        1000.00,
		Currency:      "USD",
		Counterparty: models.Counterparty{
			ID:   "", // Missing counterparty ID
			Name: "Test Corp",
			Type: "BUSINESS",
		},
		Timestamp: time.Now(),
	}

	result, err := service.ValidateTransaction(request)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, models.ValidationStatusFailed, result.Status)

	// Check that counterparty rule failed
	counterpartyRuleFailed := false
	for _, rule := range result.Rules {
		if rule.RuleID == "counterparty-check" && rule.Status == "FAILED" {
			counterpartyRuleFailed = true
			break
		}
	}
	assert.True(t, counterpartyRuleFailed, "Counterparty rule should have failed")
}

func TestValidationService_GetValidationResult(t *testing.T) {
	service := NewValidationService()

	result, err := service.GetValidationResult("test-validation-id")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "test-validation-id", result.ID)
	assert.Equal(t, models.ValidationStatusPassed, result.Status)
}