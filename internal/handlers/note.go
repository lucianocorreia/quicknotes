package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"text/template"

	"github.com/lucianocorreia/quicknotes/internal/apperror"
)

type noteHandler struct{}

func NewNoteHandler() *noteHandler {
	return &noteHandler{}
}

func (nh *noteHandler) NotesList(w http.ResponseWriter, r *http.Request) {
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

func (nh *noteHandler) NotesNew(w http.ResponseWriter, r *http.Request) {
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

func (nh *noteHandler) NotesView(w http.ResponseWriter, r *http.Request) error {
	id := r.URL.Query().Get("id")
	if id == "" {
		return apperror.WithStatus(errors.New("note ID is required"), http.StatusBadRequest)
	}

	files := []string{
		"views/templates/base.go.tmpl",
		"views/templates/pages/note-view.go.tmpl",
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		return errors.New("internal server error")
	}

	return t.ExecuteTemplate(w, "base", id)
}

func (nh *noteHandler) NotesCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Creating note")
}
