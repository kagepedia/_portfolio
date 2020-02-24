package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"server/application"
)

// TaskController : TaskにおけるControllerのインテーフェース
type TaskController interface {
	Index(http.ResponseWriter, *http.Request, httprouter.Params)
}

type taskController struct {
	taskApp application.TaskApp
}

// NewTaskApp : Taskデータに関するControllerを生成
func NewTaskController(ta application.TaskApp) TaskController {
	return &taskController{
		taskApp: ta,
	}
}

// TaskIndex : GET /task -> Taskデータ一覧を返す
func (tc taskController) Index(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	// request : task
	// -> こんな感じでリクエストも受け取れますが今回は使用しない
	type request struct {
		Begin uint `query:"begin"`
		Limit uint `query:"limit"`
	}

	// taskField : response内で使用するTaskを表す構造体
	//  -> ドメインモデルの Task に HTTP の関心事である JSON タグを付与したくないために Controller 層で用意
	//     簡略化のために JSON タグを付与したドメインモデルを流用するプロジェクトもしばしば見かける
	type taskField struct {
		Id         int64     `json:"id"`
		Name       string    `json:"name"`
		CreateTime time.Time `json:"create_at"`
		UpdateTime time.Time `json:"update_at"`
	}

	// response : task APIのレスポンス
	type response struct {
		Tasks []taskField `json:"tasks"`
	}

	ctx := r.Context()

	// Controller 呼出
	tasks, err := tc.taskApp.GetAll(ctx)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// 取得したドメインモデルをresponseに変換
	res := new(response)
	for _, task := range tasks {
		var tf taskField
		tf = taskField(*task)
		res.Tasks = append(res.Tasks, tf)
	}

	// クライアントにレスポンスを返却
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
}
