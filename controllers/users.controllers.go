package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"projet-forum/models"
	"projet-forum/services"

	"github.com/gorilla/mux"
)

type UsersControllers struct {
	service  *services.UsersServices
	template *template.Template
}

func InitUsersControllers(service *services.UsersServices, template *template.Template) *UsersControllers {
	return &UsersControllers{service: service, template: template}
}

func (c *UsersControllers) UsersRouter(r *mux.Router) {
	r.HandleFunc("/register", c.RegisterForm).Methods("GET")
	r.HandleFunc("/register/submit", c.RegisterSubmit).Methods("POST")
}

func (c *UsersControllers) RegisterForm(w http.ResponseWriter, r *http.Request) {
	c.template.ExecuteTemplate(w, "register", nil)
}

func (c *UsersControllers) RegisterSubmit(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 50)
	if err != nil {
		http.Error(w, "Erreur lors du parsing du formulaire", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	fmt.Printf("Email: %s, Password: %s\n", email, password)

	// Tenter de récupérer le fichier
	//file, handler, err := r.FormFile("image")}

	user := models.User{
		Email:    email,
		Password: password,
	}

	c.service.Create(user)

	http.Redirect(w, r, "/", http.StatusOK)
}
