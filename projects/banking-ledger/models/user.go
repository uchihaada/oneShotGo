// Package models defines database models and validations
// Learning focus: GORM model definition, struct tags, relationships, hooks
package models

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents a user in the banking system
type User struct {
	gorm.Model
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	FirstName string
	LastName  string

	// Accounts []Account `gorm:"foreignKey:UserID"` - one-to-many relationship
	Accounts []Account `gorm:"foreignKey:UserID"`
}

// BeforeCreate hook runs before creating user in database
func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return err
	}

	if err = u.ValidatePassword(u.Password); err != nil {
		log.Println(err)
		return err
	}
	u.Password = string(hashedPassword)

	return nil
}

// ValidatePassword checks if password meets requirements
func (u *User) ValidatePassword(password string) error {
	if len(password) < 6 && password != "" {
		return fmt.Errorf("password must be at least 6 characters long")
	}
	return nil
}

// CheckPassword verifies password against hashed password
func (u *User) CheckPassword(password string) error {
	// Compare provided password with u.Password (hashed)
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
