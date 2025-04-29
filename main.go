package main

import (
	"fmt"
	"log"
	"net/http"
	"projet-forum/controllers"
	"projet-forum/routes"
)

func main() {
	var err error
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
