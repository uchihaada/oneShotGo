// Package services contains business logic and concurrent processing
// Learning focus: Goroutines, channels, error handling in concurrent code
package services

import (
	"banking-ledger/database"
	"banking-ledger/models"
)

// TODO: Import database and models packages

// AccountRequest represents single account creation request
type AccountRequest struct {
	// TODO: Define fields matching CreateAccountRequest
	AccountName    string
	AccountType    string
	InitialBalance float64
}

// BulkAccountResult represents result of single account creation
type BulkAccountResult struct {
	// TODO: Define result fields
	Index       int
	Success     bool
	AccountID   uint
	AccountName string
	Error       string
}

// CreateAccountsConcurrently creates multiple accounts using goroutines
func CreateAccountsConcurrently(userID uint, requests []AccountRequest) []BulkAccountResult {
	// TODO: Create results slice
	results := make([]BulkAccountResult, len(requests))

	// TODO: Create channel for results
	resultChan := make(chan BulkAccountResult, len(requests))

	// TODO: Launch goroutine for each account
	for i, req := range requests {
		go func(index int, request AccountRequest) {
			// Create account model
			acc := models.Account{
				UserID:      userID,
				AccountName: request.AccountName,
				AccountType: models.AccountType(request.AccountType),
				Balance:     request.InitialBalance,
				IsActive:    true,
			}
			// Validate account
			if err := acc.ValidateAccount(); err != nil {
				resultChan <- BulkAccountResult{
					Index:   index,
					Success: false,
					Error:   err.Error(),
				}
				return
			}

			// Save to database
			if err := database.DB.Create(&acc).Error; err != nil {
				resultChan <- BulkAccountResult{
					Index:   index,
					Success: false,
					Error:   "failed to create account",
				}
				return
			}

			// Send result to channel
			resultChan <- BulkAccountResult{
				Index:       index,
				Success:     true,
				AccountID:   acc.ID,
				AccountName: acc.AccountName,
			}
		}(i, req)
	}

	// TODO: Collect all results from channel
	for range requests {
		result := <-resultChan
		results[result.Index] = result
	}

	// TODO: Close channel
	close(resultChan)

	return results
}
