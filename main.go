package main

import (
	"go-gin-auth/controllers"
	"go-gin-auth/initializers"
	"go-gin-auth/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

// consider putting routes in a different file subsequently, to mave the code cleaner
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello darkness my old friend!")
	})
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/user", middleware.VerifyToken, controllers.User)
	r.POST("/forgot-password", controllers.ForgotPassword)
	r.POST("/verify-otp", controllers.VerifyOtp)
	r.Run() // listen and serve on 0.0.0.0:3000
}
