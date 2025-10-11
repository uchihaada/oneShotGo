package main

import (
	"banking-ledger/database"
	"banking-ledger/handlers"
	"banking-ledger/middleware"
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

	// healthCheck
	router.GET("/health", healthCheck)

	api := router.Group("api/auth")
	{
		api.POST("/register", handlers.RegisterUser)
		api.POST("/login", handlers.LoginUser)
	}

	protected := router.Group("/protected")
	// health check
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/health", healthCheck)
	}
	//account handlers
	accounts := router.Group("/v1")
	protected.Use(middleware.AuthRequired())
	{
		accounts.GET("/accounts", handlers.GetAccounts)
		accounts.POST("/accounts", handlers.CreateAccount)
		accounts.POST("/accounts/bulk", handlers.CreateAccountsBulk)
		accounts.GET("/accounts/:id", handlers.GetAccount)
		accounts.PUT("/accounts/:id", handlers.UpdateAccount)
		accounts.DELETE("/accounts/:id", handlers.DeleteAccount)
	}
	// Transaction routes - we'll add these later
	transactions := router.Group("/v1")
	protected.Use(middleware.AuthRequired())
	{
		transactions.POST("/transactions", handlers.CreateTransaction)
		transactions.GET("/transactions", handlers.GetTransactions)
		transactions.GET("/transactions/:id", handlers.GetTransaction)
	}

	// Concurrency testing routes - we'll add these later

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
