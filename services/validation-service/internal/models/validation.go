package models

import (
	"time"
)

// ValidationRequest represents a transaction validation request
type ValidationRequest struct {
	TransactionID string                 `json:"transaction_id" binding:"required"`
	Type          string                 `json:"type" binding:"required"`
	Amount        float64                `json:"amount" binding:"required,gt=0"`
	Currency      string                 `json:"currency" binding:"required,len=3"`
	Counterparty  Counterparty           `json:"counterparty" binding:"required"`
	Metadata      map[string]interface{} `json:"metadata,omitempty"`
	Timestamp     time.Time              `json:"timestamp"`
}

// Counterparty represents transaction counterparty information
type Counterparty struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
}

// ValidationResult represents the result of a transaction validation
type ValidationResult struct {
	ID            string                 `json:"id"`
	TransactionID string                 `json:"transaction_id"`
	Status        ValidationStatus       `json:"status"`
	Rules         []RuleResult           `json:"rules"`
	ErrorCode     string                 `json:"error_code,omitempty"`
	ErrorMessage  string                 `json:"error_message,omitempty"`
	ProcessedAt   time.Time              `json:"processed_at"`
	ProcessingTime time.Duration         `json:"processing_time"`
	Metadata      map[string]interface{} `json:"metadata,omitempty"`
}

// ValidationStatus represents the validation status
type ValidationStatus string

const (
	ValidationStatusPending ValidationStatus = "PENDING"
	ValidationStatusPassed  ValidationStatus = "PASSED"
	ValidationStatusFailed  ValidationStatus = "FAILED"
	ValidationStatusError   ValidationStatus = "ERROR"
)

// RuleResult represents the result of a single validation rule
type RuleResult struct {
	RuleID      string    `json:"rule_id"`
	RuleName    string    `json:"rule_name"`
	Status      string    `json:"status"` // PASSED, FAILED, SKIPPED
	Message     string    `json:"message,omitempty"`
	ProcessedAt time.Time `json:"processed_at"`
}

// ValidationRule represents a validation rule configuration
type ValidationRule struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Type        string                 `json:"type"` // AMOUNT, CURRENCY, COUNTERPARTY, etc.
	Enabled     bool                   `json:"enabled"`
	Priority    int                    `json:"priority"`
	Config      map[string]interface{} `json:"config"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}