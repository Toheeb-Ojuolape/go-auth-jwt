package controllers

import (
	"go-gin-auth/initializers"
	"go-gin-auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchTodos(c *gin.Context) {
	var todos []models.TodoList

	result := initializers.DB.Find(&todos)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": todos,
	})
}
