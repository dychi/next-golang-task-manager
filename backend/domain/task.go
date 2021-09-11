package domain

type Task struct {
	ID       int
	Name     string
	Deadline string
	// Status string
}

// func GetTaskByID(DB *sql.DB, taskID int) (*Task, error) {
// 	// インフラストラクチャレイヤの実装を利用
// 	taskDTO, err := infrastructure.GetByTaskID(DB, taskID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	task := &Task{
// 		ID:         taskDTO.ID,
// 		Name:       taskDTO.Name,
// 		Created_at: taskDTO.Created_at,
// 	}
// 	return task, nil
// }
