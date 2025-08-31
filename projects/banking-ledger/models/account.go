// Package models - Account model definition
// Learning focus: Foreign keys, decimal handling, model validation, relationships
package models

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

// AccountType defines the type of bank account
type AccountType string

const (
	AccountTypeChecking AccountType = "checking"
	AccountTypeSavings  AccountType = "savings"
	AccountTypeBusiness AccountType = "business"
)

// Account represents a bank account
type Account struct {
	gorm.Model
	UserID      uint        `gorm:"not null;index"`
	AccountName string      `gorm:"not null"`
	AccountType AccountType `gorm:"type:varchar(20);default:'checking'"`
	Balance     float64     `gorm:"default:0;check:balance >= 0"`
	IsActive    bool        `gorm:"default:true"`

	// Relationships
	User User `gorm:"foreignKey:UserID"`

	SentTransactions     []Transaction `gorm:"foreignKey:FromAccountID"`
	ReceivedTransactions []Transaction `gorm:"foreignKey:ToAccountID"`
}

// ValidateAccount validates account data before saving
func (a *Account) ValidateAccount() error {

	//account name cannot be empty
	if a.AccountName == "" {
		log.Println("account name cannot be empty")
		return fmt.Errorf("account name cannot be empty")
	}

	// checking account type
	switch a.AccountType {
	case AccountTypeChecking, AccountTypeSavings, AccountTypeBusiness:
		// valid type
	default:
		log.Println("invalid account type")
		return fmt.Errorf("invalid account type")
	}

	// balance cannot be negative
	if a.Balance < 0 {
		log.Println("balance cannot be negative")
		return fmt.Errorf("balance cannot be negative")
	}

	return nil
}

// UpdateBalance updates account balance safely
func (a *Account) UpdateBalance(amount float64) error {

	// amount cannot be zero
	if amount == 0 {
		log.Println("amount must be non-zero")
		return errors.New("amount must be non-zero")
	}

	// withdrawals will not be possible if insufficient funds
	newBalance := a.Balance + amount
	if newBalance < 0 {
		log.Println("insufficient funds for withdrawal")
		return fmt.Errorf("insufficient funds")
	}

	// Update balance
	a.Balance = newBalance

	return nil
}

// GetBalance returns current account balance
func (a *Account) GetBalance() float64 {
	return a.Balance
}

// IsOwnedBy checks if account belongs to specific user
func (a *Account) IsOwnedBy(userID uint) bool {
	return a.UserID == userID
}
