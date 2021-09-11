package interfaces

import (
	"net/http"
	"task-manager/usecase"
)

// Task に対する Handler のインターフェース
type TaskHandler interface {
	HandleTaskGet(http.ResponseWriter, *http.Request)
	HandleTaskPost(http.ResponseWriter, *http.Request)
}

type taskHandler struct {
	taskUseCase usecase.TaskUseCase
}

// Task データに関する Handlerを生成
func NewTaskHandler(tu usecase.TaskUseCase) TaskHandler {
	return &taskHandler{
		taskUseCase: tu,
	}
}

// タスク情報を取得
func (th taskHandler) HandlerTaskGet(w http.ResponseWriter, r *http.Request) {
	// C
}
