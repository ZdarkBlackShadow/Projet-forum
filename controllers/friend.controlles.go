package controllers

import (
	"html/template"
	"net/http"
	"projet-forum/services"

	"github.com/gorilla/mux"
)

type FriendControllers struct {
	service  *services.FriendService
	template *template.Template
}

func InitFriendControllers(service *services.FriendService, template *template.Template) *FriendControllers {
	return &FriendControllers{
		service:  service,
		template: template,
	}
}

func (c *FriendControllers) FriendRouter(r *mux.Router) {
	r.HandleFunc("/create/friend-request", c.CreateFriendRequest).Methods("POST")
	r.HandleFunc("/accept/friend-request", c.AcceptFriendRequest).Methods("POST")
}

func (c *FriendControllers) CreateFriendRequest(w http.ResponseWriter, r *http.Request) {
	cookie, cookieErr := r.Cookie("token")
	if cookieErr != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	friendUsername := r.FormValue("friendUsername")

	serviceErr := c.service.CreateFriendRequest(cookie.Value, friendUsername)
	if serviceErr != nil {
		http.Error(w, serviceErr.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

func (c *FriendControllers) AcceptFriendRequest(w http.ResponseWriter, r *http.Request) {
	cookie, cookieErr := r.Cookie("token")
	if cookieErr != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	friendUsername := r.FormValue("friendUsername")

	serviceErr := c.service.AcceptFriendRequest(cookie.Value, friendUsername)
	if serviceErr != nil {
		http.Error(w, serviceErr.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}