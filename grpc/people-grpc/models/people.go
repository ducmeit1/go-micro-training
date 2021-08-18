package models

import (
	"gorm.io/gorm"
)

type People struct {
	*gorm.Model
	Name    string
	Age     int64
	Address string
}
