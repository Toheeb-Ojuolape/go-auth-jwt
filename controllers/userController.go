package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	// get the user details from the request
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
