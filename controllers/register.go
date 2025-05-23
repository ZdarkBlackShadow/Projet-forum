package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	Tmpl.ExecuteTemplate(w, "register", nil)
}

func RegisterSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := r.ParseMultipartForm(10 << 50)
	if err != nil {
		http.Error(w, "Erreur lors du parsing du formulaire", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	fmt.Printf("Email: %s, Password: %s\n", email, password)

	// Tenter de récupérer le fichier
	file, handler, err := r.FormFile("image")
	imageUploaded := false
	var imageName string
	var dstPath string

	if err == nil {
		defer file.Close()

		err = os.MkdirAll("images/users", os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		imageName = handler.Filename
		dstPath = filepath.Join("images/users", imageName)

		dst, err := os.Create(dstPath)
		if err != nil {
			http.Error(w, "Erreur de création du fichier", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, "Erreur de copie du fichier", http.StatusInternalServerError)
			return
		}
		imageUploaded = true
	}

	fmt.Printf("Email: %s, Password: %s\n", email, password)
	fmt.Println("image upload : " + fmt.Sprint(imageUploaded))
	if imageUploaded {
		fmt.Println("image name : " + imageName)
		fmt.Println("image path : " + dstPath)
	}

	//faire la requter sql ici

	http.Redirect(w, r, "/", http.StatusOK)
}
