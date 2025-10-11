// Package models - Transaction model definition
// Learning focus: Complex relationships, transaction types, validation, business logic
package models

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"gorm.io/gorm"
)

// TransactionType defines types of financial transactions
type TransactionType string

const (
	TransactionTransfer TransactionType = "transfer"
	TransactionDeposit  TransactionType = "deposit"
	TransactionWithdraw TransactionType = "withdraw"
)

// TransactionStatus defines transaction status
type TransactionStatus string

const (
	StatusPending   TransactionStatus = "pending"
	StatusCompleted TransactionStatus = "completed"
	StatusFailed    TransactionStatus = "failed"
)

// Transaction represents a financial transaction
type Transaction struct {
	gorm.Model
	FromAccountID *uint             `gorm:"index"`
	ToAccountID   *uint             `gorm:"index"`
	Amount        float64           `gorm:"not null;check:amount > 0"`
	Type          TransactionType   `gorm:"type:varchar(20);not null"`
	Status        TransactionStatus `gorm:"type:varchar(20);default:'pending'"`
	Reference     string            `gorm:"unique"`
	Description   string            `gorm:"type:text"`

	// Relationships
	FromAccount *Account `gorm:"foreignKey:FromAccountID"`
	ToAccount   *Account `gorm:"foreignKey:ToAccountID"`
}

// ValidateTransaction validates transaction data before processing
func (t *Transaction) ValidateTransaction() error {

	//amount is positive
	if t.Amount <= 0 {
		return errors.New("amount must be positive")
	}

	// Validate transaction types
	switch t.Type {
	case TransactionWithdraw:
		if t.FromAccountID == nil {
			return errors.New("account id cannot be empty")
		}
		if t.ToAccountID != nil {
			return errors.New("in case of withdrawl dest account should not be there")
		}
	case TransactionDeposit:
		if t.FromAccountID != nil {
			return errors.New("from account id should be empty")
		}
		if t.ToAccountID == nil {
			return errors.New("in case of deposit dest account should not be empty")
		}
	case TransactionTransfer:
		if t.FromAccountID == t.ToAccountID {
			return errors.New("both of the account cannot be same in case of transfering funds")
		}

		if t.FromAccountID == nil || t.ToAccountID == nil {
			return errors.New("nither of the accounts cannot be empty")
		}
	}

	return nil
}

// GenerateReference creates a unique reference number for transaction
func (t *Transaction) GenerateReference() {
	t.Reference = ("TXN-" + time.Now().Format("20060102-150406") + "-" + strconv.Itoa(rand.Intn(100000)))

}

// MarkAsCompleted updates transaction status to completed
func (t *Transaction) MarkAsCompleted() {
	t.Status = StatusCompleted
}

// MarkAsFailed updates transaction status to failed
func (t *Transaction) MarkAsFailed() {
	t.Status = StatusFailed
}
