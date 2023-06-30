package services

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	mailjet "github.com/mailjet/mailjet-apiv3-go"
)

func SendMail(title string, body string, to string, sessionId string, c *gin.Context) {
	mailjetClient := mailjet.NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "hello@miniemoney.com",
				Name:  "Go-auth",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: to,
					Name:  title,
				},
			},
			Subject:  title,
			TextPart: body,
			HTMLPart: body,
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	_, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Email sent successfully",
		"sessionId": sessionId,
	})
}
