package router

import (
	"net/http"
)

func InitRouting() {
	http.HandleFunc("/hello", helloHandle)
	http.HandleFunc("/goodbye", goodbyeHandle)
	http.HandleFunc("/", landingHandle)
}
