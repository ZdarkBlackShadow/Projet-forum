package services

import (
	"database/sql"
	"projet-forum/models/dto"
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

func (s *HomeServices) GetUser(token string) (dto.HomeModel, error) {
	userId, jwtErr := utils.VerifyJWT(token)
	if jwtErr != nil {
		return dto.HomeModel{}, jwtErr
	}

	user, userErr := s.usersRepo.GetById(userId)
	if userErr != nil {
		return dto.HomeModel{}, userErr
	}

	return dto.HomeModel{
		Userconnected:  true,
		ConnectingUser: user,
	}, nil

}
