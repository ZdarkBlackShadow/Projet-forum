package controllers

import (
	"log"
	"net/http"
	"projet-forum/models"
)

func Exemple(w http.ResponseWriter, r *http.Request) {
	var err error
	var Data models.Exemple = models.Exemple{
		Text: "Hello world!",
	}
	err = Templates.ExecuteTemplate(w, "index", Data)
	if err != nil {
		log.Fatal(err)
	}
}
