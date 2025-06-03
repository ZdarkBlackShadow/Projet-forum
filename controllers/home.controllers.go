package controllers

import (
	"html/template"
	"log"
	"net/http"
	"projet-forum/models"
	"projet-forum/services"

	"github.com/gorilla/mux"
)

type HomeControllers struct {
	service  *services.HomeServices
	template *template.Template
}

func InitHomeControllers(service *services.HomeServices, template *template.Template) *HomeControllers {
	return &HomeControllers{service: service, template: template}
}

func (c *HomeControllers) HomeRouter(r *mux.Router) {
	r.HandleFunc("/", c.Home).Methods("GET")
}

func (c *HomeControllers) Home(w http.ResponseWriter, r *http.Request) {
	var cookieValue string
	var data models.HomeModel
	cookie, err := r.Cookie("token")
	if err != nil {
		tempErr := c.template.ExecuteTemplate(w, "home", nil)
		if tempErr != nil {
			log.Fatal(tempErr)
		}
		return
	} else {
		cookieValue = cookie.Value
	}
	data, dataErr := c.service.GetUser(cookieValue)
	if dataErr != nil {
		log.Fatal(err)
		return
	}
	tempErr := c.template.ExecuteTemplate(w, "home", data)
	if tempErr != nil {
		log.Fatal(tempErr)
	}

}
