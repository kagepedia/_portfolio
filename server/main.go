package main

import (
	"fmt"

	"server/application"
	"server/infra/db"
	controller "server/interface/controller"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// 依存関係を注入（DI まではいきませんが一応注入っぽいことをしてる）
	// DI ライブラリを使えば、もっとスマートになるはず
	taskData := db.NewTaskData()
	taskApp := application.NewTaskApp(taskData)
	taskHandler := controller.NewTaskController(taskApp)

	// ルーティングの設定
	router := httprouter.New()
	router.GET("/api/v1/task", taskHandler.Index)

	// サーバ起動
	fmt.Println("========================")
	fmt.Println("Server Start >> http://localhost:9999")
	fmt.Println("========================")
	// log.Fatal(http.ListenAndServe(":9999", router))
}
