package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PeopleJob struct {
	Id         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	PeopleID   uuid.UUID `gorm:"type:uuid;not null"`
	PeopleSlut string    `gorm:"type:varchar(250);not null"`
	PeopleName string    `gorm:"type:varchar(250);not null"`
	JobID      uuid.UUID `gorm:"type:uuid;not null"`
	JobName    string    `gorm:"type:varchar(250);not null"`
	JobRanking string    `gorm:"type:varchar(250)"`
	Salary     float64   `gorm:"type:decimal"`
	Active     bool      `gorm:"type:boolean"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
