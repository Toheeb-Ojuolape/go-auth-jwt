package controllers

import (
	"go-gin-auth/initializers"
	"go-gin-auth/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
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

	//Look up the requested user
	var user models.User
	initializers.DB.First(&user, "email = ? ", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "This user does not exist in our database",
		})

		return
	}

	//Compare sent in password with the saved user's password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "The password you entered is wrong",
		})

		return
	}

	//Generate a JWT token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to authenticate user",
		})

		return
	}

	//Send the token as a cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true) // don't forget to edit your expiration date based on your use case

	//Send the token and other details to the user
	c.JSON(http.StatusOK, gin.H{
		"email":     user.Email,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"phone":     user.Phone,
		"token":     tokenString,
	})
}
