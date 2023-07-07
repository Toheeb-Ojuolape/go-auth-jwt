package controllers

import (
	"go-gin-auth/initializers"
	"go-gin-auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	// get the user id from the middleware
	id, _ := c.Get("id")

	var user models.User

	initializers.DB.First(&user, id)

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
