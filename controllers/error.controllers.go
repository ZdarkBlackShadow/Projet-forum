package controllers

import (
	"html/template"
	"net/http"
	"projet-forum/models/dto"

	"github.com/gorilla/mux"
)
// ErrorControllers provides HTTP handlers for handling error pages.
type ErrorControllers struct {
	template *template.Template
}

// InitErrorControllers initializes a new ErrorControllers instance.
//
// Parameters:
//   - template: pointer to the HTML template engine.
//
// Returns:
//   - pointer to an ErrorControllers instance.
func InitErrorControllers(template *template.Template) *ErrorControllers {
	return &ErrorControllers{
		template: template,
	}
}

// ErrorRouter configures routes for error handling.
//
// It sets a custom 404 Not Found handler and a general error page route.
func (c *ErrorControllers) ErrorRouter(r *mux.Router) {
	r.NotFoundHandler = http.HandlerFunc(c.NotFoundPage)
	r.HandleFunc("/error", c.ErrorPage).Methods("GET")
}

// NotFoundPage renders the 404 Not Found error page.
//
// This is automatically called when a non-existent route is accessed.
func (c *ErrorControllers) NotFoundPage(w http.ResponseWriter, r *http.Request) {
	data := dto.Error{
		Code:    "404",
		Message: "page not found",
	}
	err := c.template.ExecuteTemplate(w, "error", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// ErrorPage renders a generic error page using query parameters.
//
// Expected query parameters:
//   - code: error code to display.
//   - message: error message to display.
func (c *ErrorControllers) ErrorPage(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	data := dto.Error{
		Code:    params.Get("code"),
		Message: params.Get("message"),
	}
	err := c.template.ExecuteTemplate(w, "error", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}