package main

import (
	"fmt"
	"log"
	"net/http"
	"projet-forum/controllers"
	"projet-forum/database"
	"projet-forum/routes"
	"projet-forum/utils"
)

func main() {
	var err error
	if err := database.Init(); err != nil {
		log.Fatalf("Erreur Init: %v", err)
	}
	defer database.Close() // fermeture propre à la fin du programme
	err = utils.LoadEnvFile(".env")
	if err != nil {
        log.Fatal("Erreur lors du chargement du fichier .env :", err)
    }
	utils.DisplayPepper()
	err = controllers.Init()
	if err != nil {
		log.Fatalf("Error when trying to init the controllers : %o", err)
	}
	routes.InitRoutes()
	fileserver := http.FileServer(http.Dir("./public"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	fmt.Println("http://localhost:8000/")
	err = http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
