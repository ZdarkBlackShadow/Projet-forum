package routes

import (
	"net/http"
	"projet-forum/controllers"
)

func InitRoutes() {
	http.HandleFunc("/", controllers.Exemple)
	http.HandleFunc("/register", controllers.Register)
	//post routes
	http.HandleFunc("/register/submit", controllers.RegisterSubmit)
}