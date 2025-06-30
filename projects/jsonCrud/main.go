package main

import (
	"jsonCrud/controllers"
	"jsonCrud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostController)
	r.GET("/getPosts", controllers.GetPosts)
	r.GET("/getPost/:id", controllers.GetPostById)
	r.PUT("/updatePost/:id", controllers.UpdatePost)
	r.DELETE("/deletePost/:id", controllers.DeletePost)
	r.Run()

}
