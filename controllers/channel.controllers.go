package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"projet-forum/models/dto"
	"projet-forum/models/entity"
	"projet-forum/services"
	"strconv"

	"github.com/gorilla/mux"
)

// ChannelControllers handles HTTP requests related to channels.
type ChannelControllers struct {
	service  *services.ChannelService
	template *template.Template
}

// InitChannelControllers initializes a new ChannelControllers instance.
//
// Parameters:
//   - service: pointer to the ChannelService.
//   - template: pointer to the HTML template engine.
//
// Returns:
//   - pointer to a ChannelControllers instance.
func InitChannelControllers(service *services.ChannelService, template *template.Template) *ChannelControllers {
	return &ChannelControllers{
		service:  service,
		template: template,
	}
}

// ChannelRouter sets up routing for channel-related endpoints.
func (c *ChannelControllers) ChannelRouter(r *mux.Router) {
	r.HandleFunc("/channel/{id}", c.GetChannelById).Methods("GET")
	r.HandleFunc("/create/channel", c.CreateChannel).Methods("GET")
	r.HandleFunc("/create/channel/submit", c.Create).Methods("POST")
	r.HandleFunc("/delete/channel/{id}", c.Delete).Methods("POST")
	r.HandleFunc("/add/tag/{id}", c.AddTags).Methods("POST")
	r.HandleFunc("/remove/tag/{id}", c.RemoveTags).Methods("POST")
	r.HandleFunc("/create/tag/{id}", c.CreateTag).Methods("POST")
	r.HandleFunc("/create/invitation/{id}", c.CreateChannelInvitation).Methods("POST")
	r.HandleFunc("/accept/invitation/{id}", c.AcceptChannelInvitation).Methods("POST")
}

// GetChannelById handles the GET request to retrieve a channel by its ID.
func (c *ChannelControllers) GetChannelById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	cookie, cookieErr := r.Cookie("token")
	if cookieErr != nil {
		http.Error(w, "Erreur lors de la récupération du cookie", http.StatusBadRequest)
		return
	}

	channel, err := c.service.GetChannelById(userId, cookie.Value)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Erreur lors de la récupération du channel", http.StatusBadRequest)
		return
	}
	fmt.Println(channel)

	c.template.ExecuteTemplate(w, "channel", channel)
}

// Create handles the POST request to create a new channel.
func (c *ChannelControllers) Create(w http.ResponseWriter, r *http.Request) {

	token, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Erreur lors de la récupération du cookie", http.StatusBadRequest)
		return
	}
	err = r.ParseMultipartForm(10 << 50)
	if err != nil {
		http.Error(w, "Erreur lors du parsing du formulaire", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	description := r.FormValue("description")
	private := r.FormValue("status") == "private"

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Erreur lors de la récupération de l'image", http.StatusBadRequest)
		return
	}

	channelImage := entity.UserImage{
		File:    file,
		Handler: handler,
	}

	channelInfo := dto.ChannelCreation{
		Name:        name,
		Description: description,
		Private:     private,
	}
	defer file.Close()

	channelId, creationErr := c.service.CreateChannel(channelInfo, channelImage, token.Value)
	fmt.Println(creationErr)
	if creationErr != nil {
		http.Error(w, "Erreur lors de la création du channel", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/channel/"+strconv.Itoa(channelId), http.StatusSeeOther)
}

// Delete handles the POST request to delete a channel by ID.
func (c *ChannelControllers) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channelId := vars["id"]

	cookie, cookieErr := r.Cookie("token")
	if cookieErr != nil {
		http.Error(w, "Erreur lors de la récupération du cookie", http.StatusBadRequest)
		return
	}

	err := c.service.DeleteChannel(channelId, cookie.Value)
	if err != nil {
		http.Error(w, "Erreur lors de la suppression du channel", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

// AddTags handles the POST request to add tags to a channel.
func (c *ChannelControllers) AddTags(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channelId := vars["id"]

	cookie, cookieErr := r.Cookie("token")
	if cookieErr != nil {
		http.Error(w, "Erreur lors de la récupération du cookie", http.StatusBadRequest)
		return
	}

	tags := r.FormValue("tag")

	err := c.service.AddTagToChannel(channelId, tags, cookie.Value)
	if err != nil {
		http.Error(w, "Erreur lors de l'ajout du tag", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/channel/"+channelId, http.StatusSeeOther)
}

// RemoveTags handles the POST request to remove tags from a channel.
func (c *ChannelControllers) RemoveTags(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channelId := vars["id"]

	cookie, cookieErr := r.Cookie("token")
	if cookieErr != nil {
		http.Error(w, "Erreur lors de la récupération du cookie", http.StatusBadRequest)
		return
	}

	tags := r.Form["tag"]

	err := c.service.RemoveTagFromChannel(channelId, tags, cookie.Value)
	if err != nil {
		http.Error(w, "Erreur lors de la suppression du tag", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/channel/"+channelId, http.StatusSeeOther)
}

// CreateTag handles the POST request to create a new tag for a channel.
func (c *ChannelControllers) CreateTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channelId := vars["id"]

	cookie, cookieErr := r.Cookie("token")
	if cookieErr != nil {
		http.Error(w, "Erreur lors de la récupération du cookie", http.StatusBadRequest)
		return
	}

	tag := r.FormValue("tag")

	err := c.service.CreateTag(channelId, tag, cookie.Value)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Erreur lors de la création du tag", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/channel/"+channelId, http.StatusSeeOther)
}

// CreateChannelInvitation handles the POST request to create an invitation to a channel.
func (c *ChannelControllers) CreateChannelInvitation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channelId := vars["id"]

	cookie, cookieErr := r.Cookie("token")
	if cookieErr != nil {
		http.Error(w, "Erreur lors de la récupération du cookie", http.StatusBadRequest)
		return
	}

	user := r.FormValue("user")

	err := c.service.CreateChannelIvitation(cookie.Value, user, channelId)
	if err != nil {
		http.Error(w, "Erreur lors de la création de l'invitation", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/channel/"+channelId, http.StatusSeeOther)
}

// AcceptChannelInvitation handles the POST request to accept an invitation to a channel.
func (c *ChannelControllers) AcceptChannelInvitation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channelId := vars["id"]

	cookie, cookieErr := r.Cookie("token")
	if cookieErr != nil {
		http.Error(w, "Erreur lors de la récupération du cookie", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	channelName := r.FormValue("channelName")

	err := c.service.AcceptChannelInvitation(cookie.Value, username, channelName)
	if err != nil {
		http.Error(w, "Erreur lors de l'acceptation de l'invitation", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/channel/"+channelId, http.StatusSeeOther)
}

// DeclineInvitation handles the POST request to decline an invitation to a channel.
func (c *ChannelControllers) DeclineInvitation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channelId := vars["id"]

	cookie, cookieErr := r.Cookie("token")
	if cookieErr != nil {
		http.Error(w, "Erreur lors de la récupération du cookie", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	channelName := r.FormValue("channelName")

	err := c.service.DeclineChannelInvitation(cookie.Value, username, channelName)
	if err != nil {
		http.Error(w, "Erreur lors du refus de l'invitation", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/channel/"+channelId, http.StatusSeeOther)
}

func (c *ChannelControllers) CreateChannel(w http.ResponseWriter, r *http.Request) {
	err := c.template.ExecuteTemplate(w, "create-channel", nil)
	if err != nil {
		http.Error(w, "Erreur lors de l'exécution du template", http.StatusBadRequest)
		return
	}
}
