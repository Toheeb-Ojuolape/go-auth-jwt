package models

import (
	"gorm.io/gorm"
)

type Process struct {
	gorm.Model
	Email     string
	ProcessId string
	Process   string
}

// you can add an expiry to your Process model to check if the processId has expired
