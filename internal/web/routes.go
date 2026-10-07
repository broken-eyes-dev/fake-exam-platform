package web

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/exercise/", exerciseHandler)
}