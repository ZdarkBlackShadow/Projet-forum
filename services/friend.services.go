package services

import (
	"database/sql"
	"projet-forum/repository"
	"projet-forum/utils"
	"strconv"
)

type FriendService struct {
	friendRepo *repository.FriendRepository
	usersRepo  *repository.UsersRepository
}

func InitFriendServices(db *sql.DB) *FriendService {
	return &FriendService{
		friendRepo: repository.InitFriendRepository(db),
		usersRepo:  repository.InitUsersRepository(db),
	}
}

func (s *FriendService) CreateFriendRequest(token string, friendUsername string) error {
	userId, jwtErr := utils.VerifyJWT(token)
	if jwtErr != nil {
		return jwtErr
	}

	intUserId, convErr := strconv.Atoi(userId)
	if convErr != nil {
		return convErr
	}

	friendUserId, err := s.usersRepo.GetIdByUsername(friendUsername)
	if err != nil {
		return err
	}

	return s.friendRepo.AddFriendRequest(intUserId, friendUserId)
}

func (s *FriendService) AcceptFriendRequest(token string, friendUsername string) error {
	userId, jwtErr := utils.VerifyJWT(token)
	if jwtErr != nil {
		return jwtErr
	}

	intUserId, convErr := strconv.Atoi(userId)
	if convErr != nil {
		return convErr
	}

	friendId, err := s.usersRepo.GetIdByUsername(friendUsername)
	if err != nil {
		return err
	}

	friendRequest, err := s.friendRepo.GetFriendRequest(intUserId, friendId)
	if err != nil {
		return err
	}

	reqErr := s.friendRepo.AddFriend(intUserId, friendId)
	if reqErr != nil {
		return reqErr
	}

	return s.friendRepo.DeleteFriendRequest(friendRequest)
}
