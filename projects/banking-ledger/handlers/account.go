// Package handlers - Account CRUD operations
package handlers

import (
	"banking-ledger/database"
	"banking-ledger/middleware"
	"banking-ledger/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateAccountRequest represents account creation request
type CreateAccountRequest struct {
	AccountName    string  `json:"account_name" binding:"required"`
	AccountType    string  `json:"account_type"`
	InitialBalance float64 `json:"initial_balance"`
}

type UpdateAccountRequest struct {
	AccountName string `json:"account_name" binding:"required"`
	IsActive    bool   `json:"is_active"`
}

// CreateAccount handles POST /accounts
func CreateAccount(c *gin.Context) {
	userId := middleware.GetUserID(c)

	var req CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	acc := models.Account{
		UserID:      userId,
		AccountName: req.AccountName,
		AccountType: models.AccountType(req.AccountType),
		Balance:     req.InitialBalance,
		IsActive:    true,
	}

	if err := acc.ValidateAccount(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&acc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create account"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"account": acc})
}

// GetAccounts handles GET /accounts
func GetAccounts(c *gin.Context) {
	userID := middleware.GetUserID(c)

	var accounts []models.Account
	result := database.DB.
		Preload("SentTransactions").
		Preload("ReceivedTransactions").
		Preload("User").
		Where("user_id = ?", userID).
		Find(&accounts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch accounts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accounts": accounts})
}

// GetAccount handles GET /accounts/:id
func GetAccount(c *gin.Context) {
	accountIDStr := c.Param("id")
	accountID, _ := strconv.ParseUint(accountIDStr, 10, 64)

	userID := uint(middleware.GetUserID(c))

	var account models.Account

	res := database.DB.
		Preload("SentTransactions").
		Preload("ReceivedTransactions").
		Preload("User").
		Where("id = ? AND user_id = ?", accountID, userID).First(&account)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"account": account})
}

// UpdateAccount handles PUT /accounts/:id
func UpdateAccount(c *gin.Context) {
	accountIDStr := c.Param("id")
	accountID, _ := strconv.ParseUint(accountIDStr, 10, 64)

	jwtUserID := uint(middleware.GetUserID(c))

	var account models.Account
	if err := database.DB.First(&account, accountID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	// Ownership check
	if account.UserID != jwtUserID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not authorized to update this account"})
		return
	}

	var req UpdateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.AccountName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account name cannot be empty"})
		return
	}

	account.AccountName = req.AccountName
	account.IsActive = req.IsActive

	if err := database.DB.Save(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update account"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"account": account})
}

// DeleteAccount handles DELETE /accounts/:id
func DeleteAccount(c *gin.Context) {
	accountIDStr := c.Param("id")
	accountID, _ := strconv.ParseUint(accountIDStr, 10, 64)

	jwtUserID := uint(middleware.GetUserID(c))

	var account models.Account
	if err := database.DB.First(&account, accountID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		return
	}

	// Ownership check
	if account.UserID != jwtUserID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not authorized to delete this account"})
		return
	}

	// Business rule: must have zero balance
	if account.Balance != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot delete account with non-zero balance"})
		return
	}

	if err := database.DB.Delete(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete account"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted account",
		"account": account,
	})
}
