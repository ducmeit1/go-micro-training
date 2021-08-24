package repositories

import (
	"context"
	"gin-training/grpc/job-grpc/models"
	"gin-training/grpc/job-grpc/requests"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type JobRepositories interface {
	CreateJob(ctx context.Context, model *models.Job) (*models.Job, error)
	ListJob(ctx context.Context, req *requests.ListJobRequest) ([]*models.Job, error)
}

type dbmanager struct {
	*gorm.DB
}

func NewDBManager() (JobRepositories, error) {
	db, err := gorm.Open(postgres.Open("host=localhost user=admin password=admin dbname=job port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"))
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
