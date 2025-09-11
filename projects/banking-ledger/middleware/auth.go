// Package middleware contains HTTP middleware functions
// Learning focus: Middleware patterns, JWT validation, context passing
package middleware

import (
	"banking-ledger/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	// TODO: Import your utils package for JWT validation
)

// AuthRequired middleware validates JWT tokens
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Get Authorization header
		authHeader := c.GetHeader("Authorization")

		// TODO: Check if header exists and starts with "Bearer "
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			return
		}

		// TODO: Extract token from header (remove "Bearer " prefix)
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// TODO: Validate token using utils.ValidateToken()
		claims, err := utils.ValidateToken(tokenString)

		// TODO: Handle token validation errors
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// TODO: Set user information in context for handlers to use
		c.Set("userID", claims.UserID)
		c.Set("userEmail", claims.Email)

		// TODO: Continue to next handler
		c.Next()
	}
}

func GetUserID(c *gin.Context) uint {
	userID, exists := c.Get("userID")
	if !exists {
		return 0
	}
	return userID.(uint)
}

func GetUserEmail(c *gin.Context) string {
	email, exists := c.Get("userEmail")
	if !exists {
		return ""
	}
	return email.(string)
}
