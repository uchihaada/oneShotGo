package main

import (
	"banking-ledger/database"
	"banking-ledger/handlers"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Package main - Entry point for banking API
// This file sets up the Gin router and defines all API routes
// Learning focus: Gin setup, routing, HTTP methods, JSON responses

func main() {
	// TODO: Initialize database connection (we'll add this in Step 3)
	if err := database.Connect(); err != nil {
		log.Println(err)
	}
	// Router initialized
	router := gin.Default()

	// TODO: Add CORS middleware if needed for frontend (optional for now)

	// healthCheck
	router.GET("/health", healthCheck)

	// TODO: Create API v1 routes group "/api/v1"
	// This organizes all your API routes under /api/v1/

	// Public routes (no authentication needed) - we'll add these later
	// TODO: POST /api/v1/register - user registration
	// TODO: POST /api/v1/login - user login

	api := router.Group("api/auth")
	{
		api.POST("/register", handlers.RegisterUser)
		api.POST("/login", handlers.LoginUser)
	}

	// Protected routes (require JWT authentication) - we'll add these later
	// TODO: Apply JWT middleware to protected group
	// TODO: GET /api/v1/accounts - list user accounts
	// TODO: POST /api/v1/accounts - create account
	// TODO: GET /api/v1/accounts/:id - get account details
	// TODO: PUT /api/v1/accounts/:id - update account
	// TODO: DELETE /api/v1/accounts/:id - delete account

	// Transaction routes - we'll add these later
	// TODO: POST /api/v1/transactions - create transaction
	// TODO: GET /api/v1/transactions - list transactions

	// Concurrency testing routes - we'll add these later
	// TODO: POST /api/v1/accounts/bulk - bulk account creation with goroutines
	// TODO: POST /api/v1/transactions/bulk - bulk transactions with WaitGroups

	// running the server on localhost 8080
	router.Run("localhost:8080")
}

// healthCheck handles GET /health
// This is a simple endpoint to verify the server is running
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"message":   "Banking API is running",
		"timestamp": time.Now(),
	})

}
