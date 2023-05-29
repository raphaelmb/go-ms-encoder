package repositories

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/raphaelmb/go-ms-encoder/domain"
)

type JobRepostory interface {
	Insert(job *domain.Job) (*domain.Job, error)
	Find(id string) (*domain.Job, error)
	Update(job *domain.Job) (*domain.Job, error)
}

type JobRepostoryDb struct {
	Db *gorm.DB
}

func (repo *JobRepostoryDb) Insert(job *domain.Job) (*domain.Job, error) {
	err := repo.Db.Create(job).Error
	if err != nil {
		return nil, err
	}

	return job, nil
}

func (repo *JobRepostoryDb) Find(id string) (*domain.Job, error) {
	var job domain.Job
	repo.Db.Preload("Video").First(&job, "id = ?", id)

	if job.ID == "" {
		return nil, fmt.Errorf("job does not exist")
	}

	return &job, nil
}

func (repo *JobRepostoryDb) Update(job *domain.Job) (*domain.Job, error) {
	err := repo.Db.Save(&job).Error
	if err != nil {
		return nil, err
	}

	return job, err
}
