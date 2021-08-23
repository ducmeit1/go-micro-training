package repositories

import (
	"context"
	"gin-training/grpc/people-grpc/models"
	"gin-training/grpc/people-grpc/requests"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Embeded struct

type PeopleRepository interface {
	GetPeopleByID(context context.Context, id uuid.UUID) (*models.People, error)
	GetPeopleBySlut(context context.Context, slut string) (*models.People, error)
	CreatePeople(ctx context.Context, model *models.People) (*models.People, error)
	UpdatePeople(ctx context.Context, model *models.People) (*models.People, error)
	DeletePeople(ctx context.Context, id uuid.UUID) error
	ListPeoples(ctx context.Context, req *requests.ListPeopleRequest) ([]*models.People, error)
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

func (m *dbmanager) GetPeopleByID(ctx context.Context, id uuid.UUID) (*models.People, error) {
	people := models.People{}
	if err := m.Where(&models.People{Id: id}).Preload("Contacts").First(&people).Error; err != nil {
		return nil, err
	}

	return &people, nil
}

func (m *dbmanager) GetPeopleBySlut(context context.Context, slut string) (*models.People, error) {
	people := models.People{}
	if err := m.Where(&models.People{Slut: slut}).Preload("Contacts").First(&people).Error; err != nil {
		return nil, err
	}

	return &people, nil
}

func (m *dbmanager) GetContactsByPeopleID(ctx context.Context, peopleID uuid.UUID) ([]*models.Contact, error) {
	contacts := []*models.Contact{}
	if err := m.Where(&models.Contact{PeopleID: peopleID}).Preload("Contacts").Find(&contacts).Error; err != nil {
		return nil, err
	}

	return contacts, nil
}

func (m *dbmanager) CreatePeople(ctx context.Context, model *models.People) (*models.People, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (m *dbmanager) UpdatePeople(ctx context.Context, model *models.People) (*models.People, error) {
	if err := m.Where(&models.People{Id: model.Id}).Updates(model).Error; err != nil {
		return nil, err
	}

	if len(model.Contacts) > 0 {
		if err := m.Where(&models.Contact{PeopleID: model.Id}).Updates(model.Contacts).Error; err != nil {
			return nil, err
		}
	}

	return model, nil
}

func (m *dbmanager) DeletePeople(ctx context.Context, id uuid.UUID) error {
	contacts, err := m.GetContactsByPeopleID(ctx, id)
	if err != nil {
		return err
	}

	if err := m.Unscoped().Delete(&contacts).Error; err != nil {
		return err
	}

	people, err := m.GetPeopleByID(ctx, id)
	if err != nil {
		return err
	}

	return m.Unscoped().Delete(&people).Error
}

func (m *dbmanager) ListPeoples(ctx context.Context, req *requests.ListPeopleRequest) ([]*models.People, error) {
	peoples := []*models.People{}

	if req.Age > 0 {
		if err := m.Where("age >= ?", req.Age).Preload("Contacts").Find(&peoples).Error; err != nil {
			return nil, err
		}
	} else {
		if err := m.Preload("Contacts").Find(&peoples).Error; err != nil {
			return nil, err
		}
	}

	return peoples, nil
}
