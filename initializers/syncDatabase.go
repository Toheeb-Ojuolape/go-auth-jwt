package initializers

import "go-gin-auth/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})

}
