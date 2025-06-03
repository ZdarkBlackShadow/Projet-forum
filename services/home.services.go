package services

import (
	"database/sql"
	"projet-forum/models"
	"projet-forum/repository"
	"projet-forum/utils"
)

type HomeServices struct {
	usersRepo *repository.UsersRepository
}

func InitHomeServices(db *sql.DB) *HomeServices {
	return &HomeServices{
		usersRepo: repository.InitUsersRepository(db),
	}
}

func (s *HomeServices) GetUser(token string) (models.HomeModel, error) {
	userId, jwtErr := utils.VerifyJWT(token)
	if jwtErr != nil {
		return models.HomeModel{}, jwtErr
	}

	user, userErr := s.usersRepo.GetById(userId)
	if userErr != nil {
		return models.HomeModel{}, userErr
	}

	return models.HomeModel{
		Userconnected:  true,
		ConnectingUser: user,
	}, nil

}
