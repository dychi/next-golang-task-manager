package usecase

import (
	"database/sql"
	"task-manager/domain"
	"task-manager/domain/repository"
)

// Task における UseCase のインターフェース
type TaskUseCase interface {
	Insert(DB *sql.DB, name, deadline string) error
	GetByTaskID(DB *sql.DB, taskID int) (*domain.Task, error)
}

type taskUseCase struct {
	taskRepository repository.TaskRepository
}

// Task データに対する usecase を生成
func NewTaskUseCase(tr repository.TaskRepository) TaskUseCase {
	return &taskUseCase{
		taskRepository: tr,
	}
}

func (tu taskUseCase) GetByTaskID(DB *sql.DB, taskID int) (*domain.Task, error) {
	task, err := tu.taskRepository.GetByTaskID(DB, taskID)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (tu taskUseCase) Insert(DB *sql.DB, name, deadline string) error {
	// バリデーションなど
	// domain を介して infrastructure で実装した関数を呼び出す
	// Persistence (Repository) を呼出
	err := tu.taskRepository.Insert(DB, name, deadline)
	if err != nil {
		return err
	}
	return nil
}
