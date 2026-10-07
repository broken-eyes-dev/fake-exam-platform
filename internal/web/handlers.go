package web

import (
	"bytes"
	"html/template"
	"net/http"
	"strings"

	"github.com/yuin/goldmark"

	"github.com/broken-eyes-dev/fake-exam-platform/internal/exercise"
)

type PageData struct {
	Statement template.HTML
}

func renderTemplate(w http.ResponseWriter, name string, data any) {
	tmpl, err := template.ParseFiles("templates/" + name)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

		err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)

}

func exerciseHandler(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/exercise/")

	if name == ""{
		http.NotFound(w, r)
		return
	}

	path := "/exercise/level1/" + name + "/statement.md"

	statement, err := exercise.LoadStatement(path)
	if err != nil {
		http.NotFound(w, r)
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

	renderTemplate(w, "index.html", data)

}

