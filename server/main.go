package main

import (
	"fmt"
	"net/http"
	"server/infra/router"
)

func main() {
	// サーバ起動
	fmt.Println("========================")
	fmt.Println("Server Start >> http://localhost:8888")
	fmt.Println("========================")
	// ルーティング呼び出し
	router.InitRouting()
	if err := http.ListenAndServe(":8888", nil); err != nil {
		fmt.Println(err)
	}
}
