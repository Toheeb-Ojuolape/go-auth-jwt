package controllers

import (
	"fmt"
	"go-gin-auth/helpers"
	"go-gin-auth/initializers"
	"go-gin-auth/models"
	"go-gin-auth/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ForgotPassword(c *gin.Context) {
	var body struct {
		Email string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Otp not sent successfully",
		})

		return
	}

	//create a session record of the otp with its session id in postgreSQL
	//set expiry to 10 minutes
	expiry := time.Now().Add(10 * time.Minute)

	//Create the otp
	otp := models.Otp{Email: body.Email, SessionId: fmt.Sprintf("sessionId%v", helpers.GenerateOtp()), Otp: fmt.Sprint(helpers.GenerateOtp()), ExpiredAt: expiry}

	err := initializers.DB.Create(&otp)

	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Otp not sent successfully",
		})

		return
	} else {
		services.SendMail(
			"Reset your password",
			fmt.Sprintf("<h1>Hey %v </h1> <p>Sorry you forgot your password. Kindly use this otp to reset your password: <strong>%v</strong></p>", body.Email, helpers.GenerateOtp()),
			string(body.Email),
			c,
		)
	}
}
