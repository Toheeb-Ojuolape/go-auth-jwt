package services

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func SendMail(title string, body string, to string, c *gin.Context) {
	// Set up email details
	apiKey := os.Getenv("MAILGUN_APIKEY")
	domain := os.Getenv("MAILGUN_DOMAIN")
	sender := os.Getenv("MAILGUN_EMAIL")
	recipient := to
	subject := title
	html := body

	// Create the HTTP client
	client := &http.Client{}

	// Create the request payload
	data := url.Values{}
	data.Set("from", sender)
	data.Set("to", recipient)
	data.Set("subject", subject)
	data.Set("html", html)

	// Build the API URL
	apiURL := fmt.Sprintf("https://api.mailgun.net/v3/%s/messages", domain)

	// Create the HTTP POST request
	req, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	// Set the request headers
	req.SetBasicAuth("api", apiKey)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Email sending failed with status: %v", resp.Status),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Email sent successfully",
		})
	}
}
