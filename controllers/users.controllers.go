package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"projet-forum/models"
	"projet-forum/services"
	"projet-forum/utils"
	"strconv"

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
	r.HandleFunc("/connect/submit", c.Login).Methods("POST")
	r.HandleFunc("/logout", c.Logout).Methods("GET")
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

	username := r.FormValue("username")
	bio := r.FormValue("bio")
	email := r.FormValue("email")
	password := r.FormValue("password")
	user := models.User{
		Username: username,
		Bio:      bio,
		Email:    email,
		Password: password,
	}

	file, handler, err := r.FormFile("image")
	
	image := models.UserImage{
		File:    file,
		Handler: handler,
	}
	defer file.Close()

	userId, userErr := c.service.Create(user, image)
	if userErr != nil {
		http.Error(w, "Erreur lors de la création de l'utilisateur", http.StatusInternalServerError)
		fmt.Println(userErr)
		return
	}

	jwtToken, jwtErr := utils.GenerateJWT(strconv.Itoa(userId))
	if jwtErr != nil {
		http.Error(w, "Erreur lors de la génération du token JWT", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    jwtToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	http.Redirect(w, r, "/register", http.StatusSeeOther)
}

func (c *UsersControllers) Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("usernameOrEmail")
	password := r.FormValue("password")
	user, err := c.service.Connect(email, password)
	if err != nil {
		http.Error(w, "Erreur lors de la connexion de l'utilisateur", http.StatusInternalServerError)
		return
	}

	jwtToken, jwtErr := utils.GenerateJWT(strconv.Itoa(user.UserID))
	if jwtErr != nil {
		http.Error(w, "Erreur lors de la génération du token JWT", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    jwtToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (c *UsersControllers) Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
