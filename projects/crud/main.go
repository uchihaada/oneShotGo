package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var books = []Book{
	{ID: 1, Name: "ada"},
	{ID: 2, Name: "golu"},
}

func getBooks(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func createBooks(c *gin.Context) {
	var newBook Book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	books = append(books, newBook)
	c.JSON(http.StatusOK, books)
}

func updateBook(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	var newBook Book

	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	for i, x := range books {
		if x.ID == id {
			// ensure the ID remains the same even if JSON didn't include it
			newBook.ID = id
			books[i] = newBook
		}
	}

	c.JSON(http.StatusOK, books)

}
func deleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, x := range books {
		if x.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "deleted book",
			})
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"message": "book not found",
	})
}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks)
	router.POST("/books", createBooks)
	router.PUT("/books/:id", updateBook)
	router.DELETE("/books/:id", deleteBook)

	router.Run(":8080")
}
