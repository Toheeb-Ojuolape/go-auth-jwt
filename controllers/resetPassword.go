package controllers

import (
	"fmt"
	"go-gin-auth/initializers"
	"go-gin-auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ResetPassword(c *gin.Context) {
	//verify email of user by processid and get all the details about the user by their processID
	var body struct {
		ProcessId string
		Password  string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid parameters passed",
		})
		return
	}

	//fetch the process by processId
	var process models.Process
	processDb := initializers.DB.Where("process_id = ?", body.ProcessId).First(&process)

	if processDb.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ProcessId",
		})
		return
	}

	// fetch the details of the user by email
	var user models.User
	userDb := initializers.DB.Where("email = ?", process.Email).First(&user)

	if userDb.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ProcessId",
		})
		return
	}

	if process.Email != user.Email {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ProcessId",
		})
		return
	}

	if len(body.Password) < 9 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "You have entered a weak password, kindly use something stronger",
		})

		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	fmt.Printf("from user %v", user.Password)
	fmt.Printf("from request %v", body.Password)

	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "You have already used this password previously",
		})
		return
	} else if err == bcrypt.ErrMismatchedHashAndPassword {
		// that means password doesn't match which is what we want
		//hash the new password
		hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed to hash password",
			})
		}

		//set the new password
		if err := initializers.DB.Model(&user).Where("id = ?", user.ID).Update("password", string(hash)).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "hjh Something went wrong, kindly try again later",
			})
			return
		} else {

			//delete the process
			if err := initializers.DB.Delete(&process).Error; err != nil {
				c.JSON(http.StatusBadGateway, gin.H{
					"message": "Something went wrong",
				})
				return
			}

			// Password update successful
			c.JSON(http.StatusOK, gin.H{
				"message": "Password updated successfully",
			})
			return
		}
	} else {
		// An error occurred during the comparison
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong, kindly try again later",
		})
		return
	}

}
