package controllers

import (
	"go-gin-auth/initializers"
	"go-gin-auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchUsers(c *gin.Context) {
	var users []models.User

	result := initializers.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
