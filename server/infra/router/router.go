package router

import (
	"net/http"
	"server/interface/controller"
)

func InitRouting() {
	http.HandleFunc("/hello", helloHandle)
	http.HandleFunc("/goodbye", goodbyeHandle)
	http.HandleFunc("/", landingHandle)
}
