package controllers

import (
	"go-gin-auth/helpers"
	"go-gin-auth/initializers"
	"go-gin-auth/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func VerifyOtp(c *gin.Context) {
	// this should take in the otp and sessionId and return a processId
	var body struct {
		Otp       string
		SessionId string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameters passed",
		})
		return
	}

	//fetch the data of the session from the db
	var otp models.Otp
	session := initializers.DB.Where("session_id = ?", body.SessionId).First(&otp)

	if session.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid OTP",
		})
		return
	}

	// check if the otp session has expired
	if time.Now().Unix() > otp.ExpiredAt.Unix() {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "OTP session has expired",
		})
		return
	}

	//compare the otp stored with the otp passed
	if otp.Otp != body.Otp {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Otp passed",
		})
		return
	}

	//if all pass, delete the session and return a processId
	if err := initializers.DB.Delete(&otp).Error; err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "Something went wrong",
		})
		return
	} else {
		//create processId
		processId := helpers.GenerateProcessId()
		processIdRecord := models.Process{Email: otp.Email, ProcessId: processId, Process: "Auth"}

		err := initializers.DB.Create(&processIdRecord)

		if err.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error,
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message":   "OTP verified successfully",
				"processId": processId,
			})
		}
	}

}
