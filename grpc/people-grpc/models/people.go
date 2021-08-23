package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Contact struct {
	Id          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()`
	PhoneNumber string    `gorm:"type:varchar(12)"`
	Email       string    `gorm:"type:varchar(256)"`
	Fax         string    `gorm:"type:varchar(256)"`
	PeopleID    uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type People struct {
	Id        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()`
	Name      string         `gorm:"type:varchar(256)"`
	Slut      string         `gorm:"type:varchar(256);not null;unquie"`
	Age       int64          `gorm:"type:integer"`
	Address   sql.NullString `gorm:"type:varchar(256)"`
	Contacts  []*Contact     `gorm:"foreignKey:PeopleID;references:Id;constraint:OnUpdate:CASCADE:OnDelete:CASCADE"` //Delete luôn Contacts nếu People bị Delete
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
