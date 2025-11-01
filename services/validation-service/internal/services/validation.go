package services

import (
	"fmt"
	"time"

	"github.com/gtrs/validation-service/internal/models"

	"github.com/sirupsen/logrus"
)

// ValidationService handles transaction validation logic
type ValidationService struct {
	rules []models.ValidationRule
}

// NewValidationService creates a new validation service
func NewValidationService() *ValidationService {
	service := &ValidationService{
		rules: getDefaultValidationRules(),
	}

	logrus.WithField("rules_count", len(service.rules)).Info("Validation service initialized")
	return service
}

// ValidateTransaction validates a transaction against all enabled rules
func (s *ValidationService) ValidateTransaction(request *models.ValidationRequest) (*models.ValidationResult, error) {
	startTime := time.Now()

	result := &models.ValidationResult{
		ID:            fmt.Sprintf("val-%d", time.Now().UnixNano()),
		TransactionID: request.TransactionID,
		Status:        models.ValidationStatusPending,
		Rules:         make([]models.RuleResult, 0),
		ProcessedAt:   startTime,
		Metadata:      make(map[string]interface{}),
	}

	logrus.WithFields(logrus.Fields{
		"transaction_id": request.TransactionID,
		"validation_id":  result.ID,
		"amount":         request.Amount,
		"currency":       request.Currency,
	}).Info("Starting transaction validation")

	// Apply validation rules
	overallStatus := models.ValidationStatusPassed
	for _, rule := range s.rules {
		if !rule.Enabled {
			continue
		}

		ruleResult := s.applyRule(rule, request)
		result.Rules = append(result.Rules, ruleResult)

		if ruleResult.Status == "FAILED" {
			overallStatus = models.ValidationStatusFailed
		}
	}

	result.Status = overallStatus
	result.ProcessingTime = time.Since(startTime)

	// Set error details if validation failed
	if overallStatus == models.ValidationStatusFailed {
		result.ErrorCode = "VALIDATION_FAILED"
		result.ErrorMessage = "One or more validation rules failed"
	}

	logrus.WithFields(logrus.Fields{
		"transaction_id":   request.TransactionID,
		"validation_id":    result.ID,
		"status":           result.Status,
		"processing_time":  result.ProcessingTime,
		"rules_processed":  len(result.Rules),
	}).Info("Transaction validation completed")

	return result, nil
}

// GetValidationResult retrieves a validation result by ID
func (s *ValidationService) GetValidationResult(validationID string) (*models.ValidationResult, error) {
	// TODO: Implement actual storage retrieval
	// For now, return a mock result
	return &models.ValidationResult{
		ID:            validationID,
		TransactionID: "mock-transaction-id",
		Status:        models.ValidationStatusPassed,
		ProcessedAt:   time.Now(),
		ProcessingTime: 50 * time.Millisecond,
		Rules: []models.RuleResult{
			{
				RuleID:      "amount-limit",
				RuleName:    "Amount Limit Check",
				Status:      "PASSED",
				Message:     "Amount within acceptable limits",
				ProcessedAt: time.Now(),
			},
		},
	}, nil
}

// applyRule applies a single validation rule to a transaction
func (s *ValidationService) applyRule(rule models.ValidationRule, request *models.ValidationRequest) models.RuleResult {
	startTime := time.Now()

	result := models.RuleResult{
		RuleID:      rule.ID,
		RuleName:    rule.Name,
		Status:      "PASSED",
		ProcessedAt: startTime,
	}

	// Apply rule logic based on rule type
	switch rule.Type {
	case "AMOUNT_LIMIT":
		result = s.validateAmountLimit(rule, request)
	case "CURRENCY_CHECK":
		result = s.validateCurrency(rule, request)
	case "COUNTERPARTY_CHECK":
		result = s.validateCounterparty(rule, request)
	default:
		result.Status = "SKIPPED"
		result.Message = fmt.Sprintf("Unknown rule type: %s", rule.Type)
	}

	result.ProcessedAt = startTime

	logrus.WithFields(logrus.Fields{
		"rule_id":   rule.ID,
		"rule_name": rule.Name,
		"status":    result.Status,
		"message":   result.Message,
	}).Debug("Rule applied")

	return result
}

// validateAmountLimit validates transaction amount against limits
func (s *ValidationService) validateAmountLimit(rule models.ValidationRule, request *models.ValidationRequest) models.RuleResult {
	result := models.RuleResult{
		RuleID:   rule.ID,
		RuleName: rule.Name,
		Status:   "PASSED",
	}

	// Get max amount from rule config (default to 1,000,000)
	maxAmount := 1000000.0
	if limit, ok := rule.Config["max_amount"].(float64); ok {
		maxAmount = limit
	}

	if request.Amount > maxAmount {
		result.Status = "FAILED"
		result.Message = fmt.Sprintf("Amount %.2f exceeds maximum limit of %.2f", request.Amount, maxAmount)
	} else {
		result.Message = fmt.Sprintf("Amount %.2f is within limit of %.2f", request.Amount, maxAmount)
	}

	return result
}

// validateCurrency validates transaction currency
func (s *ValidationService) validateCurrency(rule models.ValidationRule, request *models.ValidationRequest) models.RuleResult {
	result := models.RuleResult{
		RuleID:   rule.ID,
		RuleName: rule.Name,
		Status:   "PASSED",
	}

	// Get allowed currencies from rule config
	allowedCurrencies := []string{"USD", "EUR", "GBP", "JPY"}
	if currencies, ok := rule.Config["allowed_currencies"].([]string); ok {
		allowedCurrencies = currencies
	}

	currencyAllowed := false
	for _, currency := range allowedCurrencies {
		if request.Currency == currency {
			currencyAllowed = true
			break
		}
	}

	if !currencyAllowed {
		result.Status = "FAILED"
		result.Message = fmt.Sprintf("Currency %s is not allowed", request.Currency)
	} else {
		result.Message = fmt.Sprintf("Currency %s is allowed", request.Currency)
	}

	return result
}

// validateCounterparty validates counterparty information
func (s *ValidationService) validateCounterparty(rule models.ValidationRule, request *models.ValidationRequest) models.RuleResult {
	result := models.RuleResult{
		RuleID:   rule.ID,
		RuleName: rule.Name,
		Status:   "PASSED",
	}

	// Basic counterparty validation
	if request.Counterparty.ID == "" {
		result.Status = "FAILED"
		result.Message = "Counterparty ID is required"
	} else if request.Counterparty.Name == "" {
		result.Status = "FAILED"
		result.Message = "Counterparty name is required"
	} else {
		result.Message = "Counterparty information is valid"
	}

	return result
}

// getDefaultValidationRules returns the default set of validation rules
func getDefaultValidationRules() []models.ValidationRule {
	return []models.ValidationRule{
		{
			ID:          "amount-limit",
			Name:        "Amount Limit Check",
			Description: "Validates transaction amount against maximum limits",
			Type:        "AMOUNT_LIMIT",
			Enabled:     true,
			Priority:    1,
			Config: map[string]interface{}{
				"max_amount": 1000000.0,
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:          "currency-check",
			Name:        "Currency Validation",
			Description: "Validates transaction currency against allowed currencies",
			Type:        "CURRENCY_CHECK",
			Enabled:     true,
			Priority:    2,
			Config: map[string]interface{}{
				"allowed_currencies": []string{"USD", "EUR", "GBP", "JPY"},
			},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:          "counterparty-check",
			Name:        "Counterparty Validation",
			Description: "Validates counterparty information completeness",
			Type:        "COUNTERPARTY_CHECK",
			Enabled:     true,
			Priority:    3,
			Config:      map[string]interface{}{},
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
}