// Package handlers - Transaction operations
// Learning focus: Database transactions, business logic, balance updates
package handlers

import (
	"banking-ledger/database"
	"banking-ledger/middleware"
	"banking-ledger/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

// CreateTransactionRequest represents transaction request
type CreateTransactionRequest struct {
	FromAccountID uint    `json:"from_account_id" binding:"required"`
	ToAccountID   uint    `json:"to_account_id" binding:"required"`
	Amount        float64 `json:"amount" binding:"required,gt=0"`
	Description   string  `json:"description"`
}

// CreateTransaction handles POST /transactions
func CreateTransaction(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var req CreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Start database transaction
	tx := database.DB.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to start transaction"})
		return
	}

	// Lock and load accounts
	var fromAccount, toAccount models.Account
	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ? AND user_id = ?", req.FromAccountID, userID).
		First(&fromAccount).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "source account not found"})
		return
	}

	if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("id = ? AND user_id = ?", req.ToAccountID, userID).
		First(&toAccount).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusNotFound, gin.H{"error": "destination account not found"})
		return
	}

	// Create transaction object and validate
	transaction := models.Transaction{
		FromAccountID: &req.FromAccountID,
		ToAccountID:   &req.ToAccountID,
		Amount:        req.Amount,
		Type:          models.TransactionTransfer,
		Status:        models.StatusPending,
		Description:   req.Description,
		FromAccount:   &fromAccount,
		ToAccount:     &toAccount,
	}

	// Generate unique reference
	transaction.GenerateReference()

	// Validate transaction
	if err := transaction.ValidateTransaction(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check sufficient balance
	if fromAccount.Balance < req.Amount {
		tx.Rollback()
		transaction.MarkAsFailed()
		c.JSON(http.StatusBadRequest, gin.H{"error": "insufficient balance"})
		return
	}

	// Update balances
	fromAccount.Balance -= req.Amount
	toAccount.Balance += req.Amount

	// Save account changes
	if err := tx.Save(&fromAccount).Error; err != nil {
		tx.Rollback()
		transaction.MarkAsFailed()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update source account"})
		return
	}

	if err := tx.Save(&toAccount).Error; err != nil {
		tx.Rollback()
		transaction.MarkAsFailed()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update destination account"})
		return
	}

	// Mark transaction as completed
	transaction.MarkAsCompleted()

	// Save transaction record
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save transaction"})
		return
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to commit transaction"})
		return
	}

	// Reload transaction with all relationships
	database.DB.
		Preload("FromAccount.User").
		Preload("ToAccount.User").
		Preload("FromAccount.SentTransactions").
		Preload("FromAccount.ReceivedTransactions").
		Preload("ToAccount.SentTransactions").
		Preload("ToAccount.ReceivedTransactions").
		First(&transaction, transaction.ID)

	c.JSON(http.StatusCreated, gin.H{
		"message":     "Transaction completed successfully",
		"transaction": transaction,
	})
}

// GetTransactions handles GET /transactions
func GetTransactions(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var transactions []models.Transaction
	result := database.DB.
		Preload("FromAccount.User").
		Preload("ToAccount.User").
		Preload("FromAccount.SentTransactions").
		Preload("FromAccount.ReceivedTransactions").
		Preload("ToAccount.SentTransactions").
		Preload("ToAccount.ReceivedTransactions").
		Where("from_account_id IN (SELECT id FROM accounts WHERE user_id = ?) OR to_account_id IN (SELECT id FROM accounts WHERE user_id = ?)",
			userID, userID).
		Find(&transactions)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch transactions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactions": transactions})
}

// GetTransaction handles GET /transactions/:id
func GetTransaction(c *gin.Context) {
	userID := middleware.GetUserID(c)

	// Parse transaction ID from URL
	transactionID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid transaction ID"})
		return
	}

	var transaction models.Transaction
	result := database.DB.
		Preload("FromAccount.User").
		Preload("ToAccount.User").
		Preload("FromAccount.SentTransactions").
		Preload("FromAccount.ReceivedTransactions").
		Preload("ToAccount.SentTransactions").
		Preload("ToAccount.ReceivedTransactions").
		Where("id = ? AND (from_account_id IN (SELECT id FROM accounts WHERE user_id = ?) OR to_account_id IN (SELECT id FROM accounts WHERE user_id = ?))",
			transactionID, userID, userID).
		First(&transaction)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "transaction not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transaction": transaction})
}
