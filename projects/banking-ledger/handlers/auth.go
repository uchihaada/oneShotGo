// Package handlers contains HTTP request handlers for authentication
// Learning focus: JSON binding, password hashing, database operations, error handling
package handlers

import (
	"banking-ledger/database"
	"banking-ledger/models"
	"banking-ledger/utils"
	"errors"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

// RegisterRequest represents user registration request body
type RegisterRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

// LoginRequest represents user login request body
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// RegisterResponse represents successful registration response
type RegisterResponse struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Message   string `json:"message"`
}

// RegisterUser handles POST /api/v1/register
// Creates a new user account with hashed password
func RegisterUser(c *gin.Context) {

	var req RegisterRequest
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = req.ValidateRequest()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check if user already exists
	var existingUser models.User
	database.DB.Where("email = ?", req.Email).First(&existingUser)

	if existingUser.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	// Create new user instance
	user := models.User{
		Email:     req.Email,
		Password:  req.Password,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	//Save user to database
	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	//response
	response := RegisterResponse{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Message:   "User registered successfully",
	}

	// Return success response
	c.JSON(http.StatusCreated, response)
}

// LoginUser handles POST /api/v1/login
// Authenticates user and returns success message (JWT in next step)
func LoginUser(c *gin.Context) {

	var req LoginRequest
	c.ShouldBindJSON(&req)

	// Find user by email
	var user models.User
	result := database.DB.Where("email = ?", req.Email).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"error": "no user found with email ",
		})
		return
	}

	// Verify password using user.CheckPassword()
	if err := user.CheckPassword(req.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "wrong password for email",
		})
		return
	}

	token, err := utils.GenerateToken(user.ID, req.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	// TODO: 4. Return success response (we'll add JWT token in Step 6)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user": gin.H{
			"id":         user.ID,
			"email":      user.Email,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
		},
	})

	log.Println("User logged in:", user.Email)
}

// ValidateRequest validates registration request data
func (r *RegisterRequest) ValidateRequest() error {

	// - Check email
	r.Email = strings.ToLower(strings.TrimSpace(r.Email))

	pattern := `^[\w\.-]+@[\w\.-]+\.\w{2,}$`
	re := regexp.MustCompile(pattern)
	if !re.MatchString(r.Email) {
		return errors.New("invalid email")
	}

	if len(r.Password) < 6 {
		return errors.New("password must be of atleast 6 characters")
	}

	r.FirstName = strings.TrimSpace(r.FirstName)
	r.LastName = strings.TrimSpace(r.LastName)

	return nil
}
