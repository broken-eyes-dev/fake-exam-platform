package main

import (
	"log"
	"net/http"

	"github.com/broken-eyes-dev/fake-exam-platform/internal/web"

)


func main() {
	web.RegisterRoutes()

	log.Println("Fake Exam running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

