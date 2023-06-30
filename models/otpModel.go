package models

import (
	"time"

	"gorm.io/gorm"
)

type Otp struct {
	gorm.Model
	Email     string
	SessionId string
	Otp       string
	ExpiredAt time.Time
}
