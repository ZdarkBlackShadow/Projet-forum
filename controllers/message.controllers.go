package controllers

import (
	"html/template"
	"net/http"
	"projet-forum/services"

	"github.com/gorilla/mux"
)

type MessageControllers struct {
	service  *services.MessageServices
	template *template.Template
}

func InitMessageControllers(service *services.MessageServices, template *template.Template) *MessageControllers {
	return &MessageControllers{
		service:  service,
		template: template,
	}
}

func (c *MessageControllers) MessageRouter(r *mux.Router) {
	r.HandleFunc("/create/message", c.CreateMessage).Methods("POST")
	r.HandleFunc("/delete/message", c.DeleteMessage).Methods("POST")
	r.HandleFunc("/update/message", c.UpdateMessage).Methods("POST")
}

func (c *MessageControllers) CreateMessage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	text := r.FormValue("textContent")
	channelId := r.FormValue("channelId")

	_, err = c.service.CreateMessage(text, channelId, cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/channel/" + channelId, http.StatusSeeOther)
}

func (c *MessageControllers) UpdateMessage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	messageId := r.FormValue("messageId")
	channelId := r.FormValue("channelId")
	text := r.FormValue("textContent")

	err = c.service.UpdateMessage(messageId, text, cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/channel/" + channelId, http.StatusSeeOther)
}

func (c *MessageControllers) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	messageId := r.FormValue("messageId")
	channelId := r.FormValue("channelId")

	err = c.service.DeleteMessage(messageId, cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/channel/" + channelId, http.StatusSeeOther)
}
