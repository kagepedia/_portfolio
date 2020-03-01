package router

import (
	"net/http"
	"server/interface/controller"
)

func InitRouting() {
	http.HandleFunc("/task", controller.Handler2)
	http.HandleFunc("/", controller.Handler)
}
