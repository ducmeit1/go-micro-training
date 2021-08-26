package responses

import (
	"time"

	"github.com/google/uuid"
)

type AccountBalance struct {
	PeopleId       uuid.UUID `gorm:"column:id"`
	AccountBalance float64   `gorm:"column:account_balance"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}
