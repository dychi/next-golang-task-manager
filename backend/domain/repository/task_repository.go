package repository

import (
	"database/sql"
	"task-manager/domain"
)

type TaskRepository interface {
	Insert(DB *sql.DB, name, deadline string) error
	GetByTaskID(DB *sql.DB, taskID int) (*domain.Task, error)
}
