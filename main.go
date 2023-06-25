package main

import (
	"go-gin-auth/controllers"
	"go-gin-auth/initializers"
	"go-gin-auth/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/user", middleware.VerifyToken, controllers.User)
	r.Run() // listen and serve on 0.0.0.0:3000
}
