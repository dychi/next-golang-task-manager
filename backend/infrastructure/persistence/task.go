package persistence

import (
	"database/sql"
	"task-manager/domain"
	"task-manager/domain/repository"
)

type taskPersistence struct{}

func NewTaskPersistence() repository.TaskRepository {
	return &taskPersistence{}
}

// タスク登録
func (tp taskPersistence) Insert(DB *sql.DB, name, deadline string) error {
	stmt, err := DB.Prepare("INSERT INTO task(name, deadline) VALUES(?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name, deadline)
	return err
}

// タスク情報を取得
func (tp taskPersistence) GetByTaskID(DB *sql.DB, taskID int) (*domain.Task, error) {
	// DB にアクセスするロジック
	row := DB.QueryRow("SELECT id, name, deadline FROM task WHERE id=?", taskID)
	// row 型をgolangで利用できる形にキャストする
	return convertToTask(row)
}

// row 型を task 型に紐付ける
func convertToTask(row *sql.Row) (*domain.Task, error) {
	task := domain.Task{}
	err := row.Scan(&task.ID, &task.Name, &task.Deadline)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &task, nil
}

// type TaskDTO struct {
// ID         int
// Name       string
// Created_at string
// }
// // タスク情報を取得
// func GetByTaskID(DB *sql.DB, taskID int) (*TaskDTO, error) {
// // DB にアクセスするロジック
// row := DB.QueryRow("SELECT id, name, created_at FROM task WHERE id=?", taskID)
// return convertToTask(row)
// }
// // row 型を taskDTO 型に紐付ける
// func convertToTask(row *sql.Row) (*TaskDTO, error) {
// taskDTO := TaskDTO{}
// err := row.Scan(&taskDTO.ID, &taskDTO.Name, &taskDTO.Created_at)
// if err != nil {
// if err == sql.ErrNoRows {
// return nil, nil
// }
// return nil, err
// }
// return &taskDTO, nil
// }
//
