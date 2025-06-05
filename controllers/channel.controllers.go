package controllers

import (
	"html/template"
	"net/http"
	"projet-forum/models/dto"
	"projet-forum/models/entity"
	"projet-forum/services"
	"strconv"

	"github.com/gorilla/mux"
)

type ChannelControllers struct {
	service  *services.ChannelService
	template *template.Template
}

func InitChannelControllers(service *services.ChannelService, template *template.Template) *ChannelControllers {
	return &ChannelControllers{
		service:  service,
		template: template,
	}
}

func (c *ChannelControllers) ChannelRouter(r *mux.Router) {
	r.HandleFunc("/channel/{id}", c.GetChannelById).Methods("GET")
	r.HandleFunc("/create/channel", c.Create).Methods("POST")
}

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
		http.Error(w, "Erreur lors de la récupération du channel", http.StatusBadRequest)
		return
	}

	c.template.ExecuteTemplate(w, "channel", channel)

}

func (c *ChannelControllers) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 50)
	if err != nil {
		http.Error(w, "Erreur lors du parsing du formulaire", http.StatusBadRequest)
		return
	}

	token, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "Erreur lors de la récupération du cookie", http.StatusBadRequest)
		return
	}
	err = r.ParseForm()
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

	channelImage := entity.UserImage {
		File: file,
		Handler: handler,
	}

	channelInfo := dto.ChannelCreation {
		Name: name,
		Description: description,
		Private: private,
	}
	defer file.Close()

	channelId, creationErr := c.service.CreateChannel(channelInfo, channelImage, token.Value)
	if creationErr != nil {
		http.Error(w, "Erreur lors de la création du channel", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/channel/"+ strconv.Itoa(channelId), http.StatusSeeOther)
}