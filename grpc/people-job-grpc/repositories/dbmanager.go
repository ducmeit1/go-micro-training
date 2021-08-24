package repositories

import (
	"context"
	"gin-training/database"
	"gin-training/grpc/people-job-grpc/models"
	"gin-training/grpc/people-job-grpc/requests"

	"gorm.io/gorm"
)

type PeopleJobRepositories interface {
	AssignPeopleJob(ctx context.Context, model *models.PeopleJob) (*models.PeopleJob, error)
	ListPeopleJob(ctx context.Context, req *requests.ListPeopleJob) ([]*models.PeopleJob, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (PeopleJobRepositories, error) {
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.PeopleJob{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db.Debug()}, nil
}

func (m *dbmanager) AssignPeopleJob(ctx context.Context, model *models.PeopleJob) (*models.PeopleJob, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (m *dbmanager) ListPeopleJob(ctx context.Context, req *requests.ListPeopleJob) ([]*models.PeopleJob, error) {
	peopleJobs := []*models.PeopleJob{}
	if req.Salary > 0 && req.Active {
		if err := m.Find(&peopleJobs, "salary > ? AND active = ?", req.Salary, req.Active).Error; err != nil {
			return nil, err
		}
		return peopleJobs, nil
	}

	if req.Salary > 0 {
		if err := m.Find(&peopleJobs, "salary > ?", req.Salary).Error; err != nil {
			return nil, err
		}
		return peopleJobs, nil
	}

	if err := m.Find(&peopleJobs).Error; err != nil {
		return nil, err
	}

	return peopleJobs, nil
}
