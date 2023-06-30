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
