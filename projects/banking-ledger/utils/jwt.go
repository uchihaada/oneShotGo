// Package utils contains utility functions for JWT token management
package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key")

type Claims struct {
	// TODO: Define UserID uint and Email string fields
	// TODO: Embed jwt.RegisteredClaims
	UserID               uint   `json:"user_id"` // Custom claim
	Email                string `json:"email"`   // Custom claim
	jwt.RegisteredClaims        // Embeds standard claims
}

func GenerateToken(userID uint, email string) (string, error) {
	// TODO: Create expiration time (24 hours)
	// TODO: Create Claims with user data and expiration
	claim := Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "user_token",
		},
	}
	// TODO: Create token with HS256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// TODO: Sign token with secret
	signedToken, err := token.SignedString(jwtSecret)
	// TODO: Return signed token string
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{},
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return jwtSecret, nil
		})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
