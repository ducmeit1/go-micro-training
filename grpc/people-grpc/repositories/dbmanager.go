package repositories

import (
	"context"
	"gin-training/grpc/people-grpc/models"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Embeded struct

type PeopleRepository interface {
	GetPeopleByID(context context.Context, id uuid.UUID) (*models.People, error)
	CreatePeople(ctx context.Context, model *models.People) (*models.People, error)
	UpdatePeople(ctx context.Context, model *models.People) (*models.People, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (PeopleRepository, error) {
	db, err := gorm.Open(postgres.Open("host=localhost user=admin password=admin dbname=people port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"))
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.People{},
		&models.Contact{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db.Debug()}, nil
}

func (m *dbmanager) GetPeopleByID(context context.Context, id uuid.UUID) (*models.People, error) {
	people := models.People{}
	if err := m.Where(&models.People{ID: id}).First(&people).Error; err != nil {
		return nil, err
	}

	return &people, nil
}

func (m *dbmanager) CreatePeople(ctx context.Context, model *models.People) (*models.People, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (m *dbmanager) UpdatePeople(ctx context.Context, model *models.People) (*models.People, error) {
	if err := m.Where(&models.People{ID: model.ID}).Updates(&models.People{Name: model.Name, Slut: model.Slut}).Error; err != nil {
		return nil, err
	}
	// if err := m.Save(model).Error; err != nil {
	// 	return nil, err
	// }

	return model, nil
}
