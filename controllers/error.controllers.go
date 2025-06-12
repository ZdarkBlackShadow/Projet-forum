package controllers

import (
	"html/template"
	"net/http"
	"projet-forum/models/dto"

	"github.com/gorilla/mux"
)

type ErrorControllers struct {
	template *template.Template
}

func InitErrorControllers(template *template.Template) *ErrorControllers {
	return &ErrorControllers{
		template: template,
	}
}

func (c *ErrorControllers) ErrorRouter(r *mux.Router) {
	r.NotFoundHandler = http.HandlerFunc(c.NotFoundPage)
	r.HandleFunc("/error", c.ErrorPage).Methods("GET")
}

func (c *ErrorControllers) NotFoundPage(w http.ResponseWriter, r *http.Request) {
	data := dto.Error{
		Code: "404",
		Message: "page not found",
	}
	err := c.template.ExecuteTemplate(w, "error", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (c *ErrorControllers) ErrorPage(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	data := dto.Error{
		Code: params.Get("code"),
		Message: params.Get("message"),
	}
	err := c.template.ExecuteTemplate(w, "error", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}