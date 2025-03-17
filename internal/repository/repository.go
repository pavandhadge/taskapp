package repository

import (
	"database/sql"

	model "github.com/pavandhadge/taskapp/internal/models"
)

type TaskStorage interface {
	GetAllTasks() ([]model.Task, error)
	CreateTask(task *model.Task) error
	ToggleTaskCompletion(taskID int) error
	DeleteTask(id int) error
	GetSingleTask(id int) (*model.Task, error)
}

type TaskRepo struct {
	TaskStorage TaskStorage
}

func NewStorage(db *sql.DB) TaskRepo {
	return TaskRepo{
		TaskStorage: &TaskStore{db}, // TaskStore now satisfies TaskStorage
	}
}
