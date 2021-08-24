package repositories

import (
	"context"
	"gin-training/database"
	"gin-training/grpc/job-grpc/models"
	"gin-training/grpc/job-grpc/requests"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobRepositories interface {
	CreateJob(ctx context.Context, model *models.Job) (*models.Job, error)
	FindJob(ctx context.Context, id uuid.UUID) (*models.Job, error)
	ListJob(ctx context.Context, req *requests.ListJobRequest) ([]*models.Job, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (JobRepositories, error) {
	db, err := database.NewGormDB()
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Job{},
	)

	if err != nil {
		return nil, err
	}

	return &dbmanager{db.Debug()}, nil
}

func (m *dbmanager) CreateJob(ctx context.Context, model *models.Job) (*models.Job, error) {
	if err := m.Create(model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (m *dbmanager) FindJob(ctx context.Context, id uuid.UUID) (*models.Job, error) {
	job := models.Job{}
	if err := m.First(&job, "id = ?", id.String()).Error; err != nil {
		return nil, err
	}

	return &job, nil
}

func (m *dbmanager) ListJob(ctx context.Context, req *requests.ListJobRequest) ([]*models.Job, error) {
	jobs := []*models.Job{}
	if req.Level > 0 {
		if err := m.Where("level > ?", req.Level).Find(&jobs).Error; err != nil {
			return nil, err
		}
	} else {
		if err := m.Find(&jobs).Error; err != nil {
			return nil, err
		}
	}

	return jobs, nil
}
