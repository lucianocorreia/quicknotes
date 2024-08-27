package handlers

import (
	"errors"
	"net/http"
	"text/template"

	"github.com/lucianocorreia/quicknotes/internal/apperror"
)

type HandlerWithError func(w http.ResponseWriter, r *http.Request) error

// ServeHTTP handles the HTTP request by calling the underlying handler function and handling any errors that occur.
// If an error occurs during the execution of the handler function, it will be returned as an HTTP internal server error.
// The function signature is:
//
//	func (fn HandlerWithError) ServeHTTP(w http.ResponseWriter, r *http.Request)
func (fn HandlerWithError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		var statusError *apperror.StatusError
		if ok := errors.As(err, &statusError); ok {

			if statusError.StatusCode() == http.StatusNotFound {
				files := []string{
					"views/templates/base.go.tmpl",
					"views/templates/pages/404.go.tmpl",
				}
				t, err := template.ParseFiles(files...)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				t.ExecuteTemplate(w, "base", statusError.Error())
				return
			}

			http.Error(w, statusError.Error(), statusError.StatusCode())
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
