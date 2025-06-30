package controllers

import (
	"jsonCrud/initializers"
	"jsonCrud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostController(c *gin.Context) {

	// get data from the request body
	var request struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	if c.Bind(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// create a post
	post := models.Post{Title: request.Title, Body: request.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create post",
		})
		return
	}

	//return it
	c.JSON(200, gin.H{
		"Message": "Post created successfully",
	})
}

func GetPosts(c *gin.Context) {

	// get the posts
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to retrieve posts",
		})
		return
	}

	// Response with the posts
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})

}

func GetPostById(c *gin.Context) {

	var post models.Post
	id := c.Param("id")

	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Post with id : " + id + " not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})

}

func UpdatePost(c *gin.Context) {

	// get the id from the URL
	id := c.Param("id")

	// get the post from the database
	var post models.Post
	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Post with id : " + id + " not found",
		})
		return
	}
	// get the data from the request body

	var request struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	if c.Bind(&request) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}
	// update the post
	initializers.DB.Model(&post).Updates(models.Post{Title: request.Title, Body: request.Body})

	// return the updated post
	c.JSON(http.StatusOK, gin.H{
		"Message": "Post updated successfully",
		"post":    post,
	})

}

func DeletePost(c *gin.Context) {

	// get post by id
	id := c.Param("id")
	var post models.Post
	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Post with id : " + id + " not found",
		})
		return
	}
	// delte the post
	initializers.DB.Delete(&models.Post{}, id)
	// return success message
	c.JSON(http.StatusBadGateway, gin.H{
		"Message": "Post with id : " + id + " deleted successfully",
		"post":    post,
	})

}
