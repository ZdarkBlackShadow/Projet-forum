package services

import (
	"database/sql"
	"fmt"
	"projet-forum/models"
	"projet-forum/repository"
)

type UsersServices struct {
	usersRepo *repository.UsersRepository
}

func InitUsersServices(db *sql.DB) *UsersServices {
	return &UsersServices{repository.InitUsersRepository(db)}
}

func (s *UsersServices) Create(user models.User) (int, error) {
	if user.Username == "" || user.Password == "" {
		return -1, fmt.Errorf("Erreur ajout user - Données manquantes ou invalides")
	}

	/* 	if err == nil {
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
	} */

	//à completer

	userId, userErr := s.usersRepo.CreateUser(user)
	if userErr != nil {
		return -1, userErr
	}
	return userId, nil
}
