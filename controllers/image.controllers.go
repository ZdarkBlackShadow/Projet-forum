package controllers

import (
	"html/template"
	"net/http"
	"projet-forum/services"

	"github.com/gorilla/mux"
)

type ImageControllers struct {
	service  *services.ImageServices
	template *template.Template
}

func InitImageControllers(service *services.ImageServices, template *template.Template) *ImageControllers {
	return &ImageControllers{service: service, template: template}
}

func (c *ImageControllers) ImageRouter(r *mux.Router) {
	r.HandleFunc("/image/{filename}", c.GetImageByFilename).Methods("GET")
}

func (c *ImageControllers) GetImageByFilename(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	data, content, err := c.service.GetImageByName(cookie.Value, filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", content)
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
