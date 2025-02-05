package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/likesense/task-service/internal/models"
)

type Task interface {
	GetById(id uint64) (*models.Task, error)
	GetAll() (tasks []*models.Task, err error)
	GetAllThemes() ([]string, error)
	GetByFilterList(filters ...func(any) any) (tasks []*models.Task, err error)
}
type Hint interface {
	GetByID(ID uint64) (*models.Hint, error)
	GetAllByTaskID(taskID uint64) (hints []*models.Hint, err error)
}

type Repositories struct {
	Task Task
	Hint Hint
}

func NewRepositories(postgresDb *sqlx.DB) *Repositories {
	return &Repositories{
		Task: NewTaskRepository(postgresDb),
		Hint: NewHintRepository(postgresDb),
	}
}
