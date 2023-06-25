package controllers

import (
	"go-gin-auth/initializers"
	"go-gin-auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	//Get the email/password off the req body

	var body struct {
		Email     string
		Password  string
		Username  string
		FirstName string
		LastName  string
		Phone     string
	}

	// this binds the req to the body struct
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	//Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
	}

	//Create the user
	user := models.User{Email: body.Email, Password: string(hash), Username: body.Username, FirstName: body.FirstName, LastName: body.LastName, Phone: body.Phone}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user, user already exists",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{})

}
