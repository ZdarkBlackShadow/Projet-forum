package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"projet-forum/controllers"
	"projet-forum/database"
	"projet-forum/services"
	"projet-forum/utils"

	"github.com/gorilla/mux"
)

func main() {
	utils.LoadEnv()

	db, dbErr := database.Init()

	if dbErr != nil {
		log.Fatal(dbErr.Error())
		return
	}

	defer db.Close()

	temp, tempErr := template.ParseGlob("views/*.html")

	fmt.Println("chargement des templates réussie")

	if tempErr != nil {
		log.Fatalf("Erreur chargement des templates - %s", tempErr.Error())
	}

	//Initialisation des différents services
	usersService := services.InitUsersServices(db)

	fmt.Println("Initialisation des services réussi")

	//Initialisation des différents controllers
	usersController := controllers.InitUsersControllers(usersService, temp)

	fmt.Println("Initialisation des différents controllers réussi")

	//chargement des différents routers
	router := mux.NewRouter()
	usersController.UsersRouter(router)

	staticFileDirectory := http.Dir("./public/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	router.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")


	fmt.Println("http://localhost:8000/")

	serveErr := http.ListenAndServe("localhost:8000", router)
	if serveErr != nil {
		log.Fatal(serveErr)
	}
}
