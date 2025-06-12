package controllers

import (
	"html/template"
	"log"
	"net/http"
	"projet-forum/models/dto"
	"projet-forum/services"

	"github.com/gorilla/mux"
)
// HomeControllers provides HTTP handlers for rendering the home page.
type HomeControllers struct {
	service  *services.HomeServices
	template *template.Template
}

// InitHomeControllers initializes a new HomeControllers instance.
//
// Parameters:
//   - service: pointer to the HomeServices providing home-related business logic.
//   - template: pointer to the HTML template engine.
//
// Returns:
//   - pointer to an initialized HomeControllers instance.
func InitHomeControllers(service *services.HomeServices, template *template.Template) *HomeControllers {
	return &HomeControllers{service: service, template: template}
}

// HomeRouter registers the route for the home page.
//
// Parameters:
//   - r: the Gorilla Mux router to attach the home route to.
func (c *HomeControllers) HomeRouter(r *mux.Router) {
	r.HandleFunc("/", c.Home).Methods("GET")
}

// Home renders the home page, checking if the user is authenticated.
//
// If a valid "token" cookie is found, it loads personalized home data.
// Otherwise, it renders the home page as a guest.
//
// Template used: "homeTest".
func (c *HomeControllers) Home(w http.ResponseWriter, r *http.Request) {
	var cookieValue string
	var data dto.HomeModel

	cookie, err := r.Cookie("token")
	if err != nil {
		// No valid cookie, render home as unauthenticated user
		data.Userconnected = false
		tempErr := c.template.ExecuteTemplate(w, "homeTest", data)
		if tempErr != nil {
			log.Fatal(tempErr)
		}
		return
	}

	cookieValue = cookie.Value
	data, dataErr := c.service.Home(cookieValue)
	if dataErr != nil {
		log.Fatal(dataErr)
		return
	}

	tempErr := c.template.ExecuteTemplate(w, "homeTest", data)
	if tempErr != nil {
		log.Fatal(tempErr)
	}
}