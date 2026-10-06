package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"

	"github.com/yuin/goldmark"

	"github.com/broken-eyes-dev/fake-exam-platform/internal/exercise"
)

type PageData struct {
	Statement template.HTML
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	statement, err := exercise.LoadStatement("exercises/level1/only1/statement.md")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer

	err = goldmark.Convert([]byte(statement), &buf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{
		Statement: template.HTML(buf.String()),
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", homeHandler)

	log.Println("Fake Exam running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

