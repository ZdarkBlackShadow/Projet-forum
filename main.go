package main

import (
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

	if tempErr != nil {
		log.Fatalf("Erreur chargement des templates - %s", tempErr.Error())
	}

	//Initialisation des différents services
	usersService := services.InitUsersServices(db)
	homeService := services.InitHomeServices(db)
	imageService := services.InitImageServices(db)

	//Initialisation des différents controllers
	usersController := controllers.InitUsersControllers(usersService, temp)
	homeController := controllers.InitHomeControllers(homeService, temp)
	imageController := controllers.InitImageControllers(imageService, temp)

	//chargement des différents routers
	router := mux.NewRouter()
	usersController.UsersRouter(router)
	homeController.HomeRouter(router)
	imageController.ImageRouter(router)

	//ajout du ficher public
	staticFileDirectory := http.Dir("./public/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	router.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	log.Println("http://localhost:8000/")
	serveErr := http.ListenAndServe("localhost:8000", router)
	if serveErr != nil {
		log.Fatal(serveErr)
	}
}
