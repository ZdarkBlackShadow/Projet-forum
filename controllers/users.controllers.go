// Package controllers handles HTTP request handling and routing for the application.
package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"projet-forum/models/entity"
	"projet-forum/services"
	"projet-forum/utils"
	"strconv"

	"github.com/gorilla/mux"
)

// UsersControllers handles all user-related HTTP endpoints including registration,
// authentication, profile management and user information retrieval.
type UsersControllers struct {
	service  *services.UsersServices
	template *template.Template
}

// InitUsersControllers creates a new UsersControllers instance with the provided services and template.
// It initializes the controller with necessary dependencies for handling user-related requests.
func InitUsersControllers(service *services.UsersServices, template *template.Template) *UsersControllers {
	return &UsersControllers{service: service, template: template}
}

// UsersRouter sets up all the routes related to user functionality in the application.
// It maps HTTP endpoints to their corresponding handler functions for user operations
// such as registration, login, logout, and profile management.
func (c *UsersControllers) UsersRouter(r *mux.Router) {
	r.HandleFunc("/register", c.RegisterForm).Methods("GET")
	r.HandleFunc("/register/submit", c.RegisterSubmit).Methods("POST")
	r.HandleFunc("/connect/submit", c.Login).Methods("POST")
	r.HandleFunc("/logout", c.Logout).Methods("GET")
	r.HandleFunc("/profile/{username}", c.Profile).Methods("GET")
	r.HandleFunc("/login", c.LoginGet).Methods("GET")
	r.HandleFunc("/user", c.GetAllInformationAboutTheUserWhoAreConnected).Methods("GET")
}

// RegisterForm handles GET requests to the registration page.
// It renders the registration form template for new users to sign up.
func (c *UsersControllers) RegisterForm(w http.ResponseWriter, r *http.Request) {
	c.template.ExecuteTemplate(w, "register", nil)
}

// RegisterSubmit handles POST requests for user registration.
// It processes the registration form data including user details and profile image,
// creates a new user account, and sets up authentication via JWT token.
// On success, redirects to the registration confirmation page.
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
	user := entity.User{
		Username: username,
		Bio:      bio,
		Email:    email,
		Password: password,
	}

	file, handler, err := r.FormFile("image")
	var image entity.UserImage
	if err == nil {
		image = entity.UserImage{
			File:    file,
			Handler: handler,
		}
		defer file.Close()
	}

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

// Login handles POST requests for user authentication.
// It validates the provided credentials, generates a JWT token on successful authentication,
// and sets it as an HTTP-only cookie. On success, redirects to the home page.
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

// Logout handles user sign-out requests.
// It invalidates the authentication cookie by setting an empty value and immediate expiration,
// then redirects to the home page.
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

// Profile handles requests to view a user's profile page.
// It retrieves and displays user information based on the username provided in the URL.
func (c *UsersControllers) Profile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]
	user, err := c.service.GetUser(username)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des utilisateurs", http.StatusInternalServerError)
		return
	}

	c.template.ExecuteTemplate(w, "users", user)
}

// LoginGet handles GET requests to the login page.
// It renders the login form template for user authentication.
func (c *UsersControllers) LoginGet(w http.ResponseWriter, r *http.Request) {
	c.template.ExecuteTemplate(w, "login", nil)
}

// GetAllInformationAboutTheUserWhoAreConnected handles requests to view the current user's complete profile.
// It retrieves detailed information about the authenticated user using their JWT token.
// If the user is not authenticated, they are redirected to the login page.
func (c *UsersControllers) GetAllInformationAboutTheUserWhoAreConnected(w http.ResponseWriter, r *http.Request) {
	token, cookieErr := r.Cookie("token")
	if cookieErr != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	data, err := c.service.GetAllInformationAboutOneUser(token.Value)
	if err != nil {
		http.Redirect(w, r, "/error?code=403&message=InternalServerError", http.StatusSeeOther)
	}
	c.template.ExecuteTemplate(w, "profile", data)
}
