package router

import (
	"net/http"
	"server/interface/controller"
)

// var Tc controller.TaskController
// コントローラのポインタ
var (
	Tc = &controller.TaskController{}
)

func InitRouting() {
	http.Handle("/task", Tc)
	http.HandleFunc("/", controller.Handler)
}
