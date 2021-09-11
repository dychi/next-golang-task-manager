package interfaces

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"task-manager/config"
	"task-manager/usecase"

	"github.com/julienschmidt/httprouter"
)

// Task に対する Handler のインターフェース
type TaskHandler interface {
	HandleTaskGet(http.ResponseWriter, *http.Request, httprouter.Params)
	HandleTaskPost(http.ResponseWriter, *http.Request, httprouter.Params)
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
func (th taskHandler) HandleTaskGet(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	// C

	type taskField struct {
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Deadline string `json:"deadline"`
	}
	type response struct {
		Tasks []taskField `json:"tasks"`
	}
	// Context
	id := pr.ByName("id")
	taskID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Not Found", 404)
		return
	}

	// connect DB
	db, err := config.ConnectDB()
	if err != nil {
		http.Error(w, "Database Connection Error.", 500)
		return
	}
	defer db.Close()

	// ユースケースの呼び出し
	task, err := th.taskUseCase.GetByTaskID(db, taskID)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	if task == nil {
		http.Error(w, "Invalid Request Parameter", 502)
		return
	}
	// 取得したドメインモデルを response に変換
	// res := new(response)
	log.Println(*task)
	res := taskField(*task)

	// クライアントにレスポンスを返却
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		// TODO: エラーハンドリング処理追加
		http.Error(w, "Internal Server Error.", 500)
		return
	}
}

// タスク新規登録
func (th taskHandler) HandleTaskPost(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	// リクエストボディを取得
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Invalid Request Body", 502)
		return
	}
	// connect DB
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, "Database Connection Error", 500)
		return
	}
	defer db.Close()

	type taskCreateRequest struct {
		Name     string
		Deadline string
	}
	// リクエストボディのパース
	var requestBody taskCreateRequest
	json.Unmarshal(body, &requestBody)
	// usecase の呼び出し
	err = th.taskUseCase.Insert(db, requestBody.Name, requestBody.Deadline)
	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	type postResponse struct {
		Status  int
		Message string
	}
	resBuf := postResponse{http.StatusOK, "Success"}
	resJson, err := json.Marshal(resBuf)

	// クライアントにレスポンスを返却
	w.Header().Set("Content-Type", "application/json")
	w.Write(resJson)
}
