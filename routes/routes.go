package routes

import (
	"net/http"
	"projet-forum/controllers"
)

func InitRoutes() {
	http.HandleFunc("/", controllers.Exemple)
}