package repositories

import (
	"context"
	"gin-training/grpc/people-grpc/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Embeded struct

type PeopleRepository interface {
	CreatePeople(ctx context.Context, model *models.People) (*models.People, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (PeopleRepository, error) {
	db, err := gorm.Open(postgres.Open("host=localhost user=admin password=admin dbname=people port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.People{})

	return &dbmanager{db}, nil
}

func (m *dbmanager) CreatePeople(ctx context.Context, model *models.People) (*models.People, error) {
	if err := m.Save(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}
