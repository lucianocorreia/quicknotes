package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	config := loadConfig()

	mux := http.NewServeMux()

	staticHandler := http.FileServer(http.Dir("views/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	mux.HandleFunc("/", notesList)
	mux.HandleFunc("/note/view", notesView)
	mux.HandleFunc("/note/new", notesNew)
	mux.HandleFunc("/note/create", notesCreate)

	fmt.Printf("Starting server on port %s\n", config.ServerPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), mux)
	if err != nil {
		panic(err)
	}
}

func notesList(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"views/templates/base.go.tmpl",
		"views/templates/pages/home.go.tmpl",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.ExecuteTemplate(w, "base", nil)
}

func notesNew(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"vews/templates/base.go.tmpl",
		"views/templates/pages/note-new.go.tmpl",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.ExecuteTemplate(w, "base", nil)
}

func notesView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Note ID is required", http.StatusBadRequest)
		return
	}

	files := []string{
		"views/templates/base.go.tmpl",
		"views/templates/pages/note-view.go.tmpl",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.ExecuteTemplate(w, "base", id)
}

func notesCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Creating note")
}
