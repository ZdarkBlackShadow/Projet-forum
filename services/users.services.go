package services

import (
	"database/sql"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"projet-forum/models/entity"
	"projet-forum/repository"
	"projet-forum/utils"
	"time"
)

type UsersServices struct {
	usersRepo  *repository.UsersRepository
	imagesRepo *repository.ImageRepository
}

func InitUsersServices(db *sql.DB) *UsersServices {
	return &UsersServices{
		usersRepo:  repository.InitUsersRepository(db),
		imagesRepo: repository.InitImageRepository(db),
	}
}

func (s *UsersServices) Create(user entity.User, image entity.UserImage) (int, error) {
	if user.Username == "" || user.Password == "" {
		return -1, fmt.Errorf("Erreur ajout user - Données manquantes ou invalides")
	}

	if image.File != nil && image.Handler != nil {
		err := os.MkdirAll("images/users", os.ModePerm)
		if err != nil {
			return -1, err
		}
		imageName := image.Handler.Filename
		dstPath := filepath.Join("images/users", imageName)

		dst, dstErr := os.Create(dstPath)
		if dstErr != nil {
			return -1, dstErr
		}
		defer dst.Close()

		_, copyErr := io.Copy(dst, image.File)
		if copyErr != nil {
			return -1, copyErr
		}

		imageId, imageErr := s.imagesRepo.Create(dstPath)
		if imageErr != nil {
			return -1, imageErr
		}
		user.ImageID = imageId
	} else {
		user.ImageID = 14
	}

	hashedPassword, salt, passErr := utils.HashPassword(user.Password)
	if passErr != nil {
		return -1, passErr
	}

	user.Password = hashedPassword
	user.Salt = salt

	user.LastConnection = time.Now()

	userId, userErr := s.usersRepo.Create(user)
	if userErr != nil {
		return -1, userErr
	}
	return userId, nil
}

func (s *UsersServices) Connect(nameOrMail string, password string) (entity.User, error) {
	userSalt, saltErr := s.usersRepo.GetSaltByEmailOrUsername(nameOrMail)
	if saltErr != nil {
		return entity.User{}, saltErr
	}

	hashedPassword, passErr := utils.HashPasswordWithSalt(password, userSalt)
	if passErr != nil {
		return entity.User{}, passErr
	}
	user, userErr := s.usersRepo.GetUserByEmailOrNameAndPassword(nameOrMail, hashedPassword)
	if userErr != nil {
		return entity.User{}, userErr
	}

	return user, nil
}
