package controllers

import (
	"html/template"
	"net/http"
	"projet-forum/services"

	"github.com/gorilla/mux"
)

// MessageControllers provides HTTP handlers for creating, updating, deleting,
// and voting on messages within a channel.
type MessageControllers struct {
	service  *services.MessageServices
	template *template.Template
}

// InitMessageControllers initializes a new instance of MessageControllers.
//
// Parameters:
//   - service: pointer to MessageServices which contains message-related logic.
//   - template: pointer to the template engine (not directly used here).
//
// Returns:
//   - pointer to an initialized MessageControllers instance.
func InitMessageControllers(service *services.MessageServices, template *template.Template) *MessageControllers {
	return &MessageControllers{
		service:  service,
		template: template,
	}
}

// MessageRouter registers routes for message-related actions.
//
// Parameters:
//   - r: the Gorilla Mux router to register the message routes with.
func (c *MessageControllers) MessageRouter(r *mux.Router) {
	r.HandleFunc("/create/message/{id}", c.CreateMessage).Methods("POST")
	r.HandleFunc("/delete/message", c.DeleteMessage).Methods("POST")
	r.HandleFunc("/update/message", c.UpdateMessage).Methods("POST")
	r.HandleFunc("/create/updownvote", c.AddUpDownVote).Methods("POST")
	r.HandleFunc("/update/updownvote", c.UpdateUpDownVote).Methods("POST")
}

// CreateMessage handles the creation of a new message.
//
// Expects "textContent" and "channelId" form values and a "token" cookie for authentication.
// Redirects to the channel page after successful creation.
func (c *MessageControllers) CreateMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channelId := vars["id"]
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	text := r.FormValue("textContent")

	_, err = c.service.CreateMessage(text, channelId, cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/channel/"+channelId, http.StatusSeeOther)
}

// UpdateMessage modifies the content of an existing message.
//
// Expects "messageId", "channelId", and "textContent" form values and a "token" cookie.
// Redirects to the channel page after update.
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
	http.Redirect(w, r, "/channel/"+channelId, http.StatusSeeOther)
}

// DeleteMessage removes an existing message.
//
// Expects "messageId" and "channelId" form values and a "token" cookie.
// Redirects to the channel page after deletion.
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
	http.Redirect(w, r, "/channel/"+channelId, http.StatusSeeOther)
}

// AddUpDownVote registers a new upvote or downvote for a message.
//
// Expects "messageId", "channelId", and "vote" form values and a "token" cookie.
// Redirects to the channel page after voting.
func (c *MessageControllers) AddUpDownVote(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	messageId := r.FormValue("messageId")
	channelId := r.FormValue("channelId")
	vote := r.FormValue("vote")

	err = c.service.AddUpDownVote(messageId, vote, cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/channel/"+channelId, http.StatusSeeOther)
}

// UpdateUpDownVote modifies an existing vote (upvote or downvote) on a message.
//
// Expects "messageId", "channelId", and "vote" form values and a "token" cookie.
// Redirects to the channel page after the vote update.
func (c *MessageControllers) UpdateUpDownVote(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	messageId := r.FormValue("messageId")
	channelId := r.FormValue("channelId")
	vote := r.FormValue("vote")

	err = c.service.UpdateUpDownVote(messageId, vote, cookie.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/channel/"+channelId, http.StatusSeeOther)
}