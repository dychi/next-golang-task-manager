package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	_ "github.com/go-sql-driver/mysql"
)

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello, %s!\n", ps.ByName("name"))
}

func main() {
	// 依存関係を注入
	// taskPersistence := persistence.NewTaskPersistence()
	// taskUseCase := usecase.NewTaskUseCase(taskPersistence)
	// taskHandler := interfaces.NewTaskHandler(taskUseCase)

	// ルーティングの設定
	router := httprouter.New()
	router.GET("/api/hello/:name", Hello)
	// router.GET("/api/task", taskHandler.HandleTaskGet)
	// router.POST("/api/task", taskHandler.HandleTaskPost)

	// サーバーの起動
	fmt.Println("===================")
	fmt.Println("Server Start >> http://localhost:8080")
	fmt.Println("===================")
	log.Fatal(http.ListenAndServe(":8080", router))
}
