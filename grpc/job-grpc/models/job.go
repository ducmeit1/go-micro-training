package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Job struct {
	Id          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name        string         `gorm:"type:varchar(250);not null"`
	Description sql.NullString `gorm:"type:varchar(250)"`
	Level       int64          `gorm:"type:integer"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
