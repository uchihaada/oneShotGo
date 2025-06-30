package main

import (
	"jwtgo/controllers"
	"jwtgo/initializers"
	"jwtgo/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)
	router.Run()
}
