package repository

import (
	"database/sql"
	"projet-forum/models/entity"
	"time"
)

type FriendRepository struct {
	db *sql.DB
}

func InitFriendRepository(db *sql.DB) *FriendRepository {
	return &FriendRepository{db}
}

func (r *FriendRepository) AddFriendRequest(user_id_who_invite int, user_id_who_is_invited int) error {
	query := "INSERT INTO friend_request (user_id_creator, user_id_invite, created_at) VALUES ('?', '?', '?')"

	_, resultErr := r.db.Exec(query, user_id_who_invite, user_id_who_is_invited, time.Now())
	if resultErr != nil {
		return resultErr
	}
	return nil
}

func (r *FriendRepository) GetFriendRequest(userIdCreator int, userIdInvite int) (entity.FriendRequest, error) {
	query := "SELECT user_id_creator, user_id_invite FROM friend_request WHERE user_id_creator = '?' AND user_id_invite = '?'"
	var friendRequest entity.FriendRequest
	resultErr := r.db.QueryRow(query, userIdCreator, userIdInvite).Scan(
		&friendRequest.UserIdCreator,
		&friendRequest.UserIdInvite,
	)
	if resultErr != nil {
		return entity.FriendRequest{}, resultErr
	}
	return friendRequest, nil
}

func (r *FriendRepository) DeleteFriendRequest(request entity.FriendRequest) error {
	query := "DELETE FROM friend_request WHERE user_id_creator = '?' AND user_id_invite = '?'"
	_, resultErr := r.db.Exec(query, request.UserIdCreator, request.UserIdInvite)
	if resultErr != nil {
		return resultErr
	}
	return nil
}

func (r *FriendRepository) AddFriend(userId1 int, userId2 int) error {
	query := "INSERT INTO friend (user_id_1, user_id_2, created_at) VALUES ('?', '?', '?')"

	_, resultErr := r.db.Exec(query, userId1, userId2, time.Now())
	if resultErr != nil {
		return resultErr
	}
	return nil
}

func (r *FriendRepository) GetFriend(friend entity.Friend) (entity.Friend, error) {
	query := "SELECT user_id_1, user_id_2 FROM friend WHERE (user_id_1 = '?' AND user_id_2 = '?') OR (user_id_1 = '?' AND user_id_2 = '?');"
	var result entity.Friend
	resultErr := r.db.QueryRow(query, friend.UserID, friend.UserID1, friend.UserID1, friend.UserID).Scan(
		&result.UserID,
		&result.UserID1,
	)
	if resultErr != nil {
		return entity.Friend{}, resultErr
	}
	return result, nil
}

func (r *FriendRepository) DeleteFriend(friend entity.Friend) error {
	query := "DELETE FROM friend WHERE (user_id_1 = '?' AND user_id_2 = '?') OR (user_id_1 = '?' AND user_id_2 = '?');"
	_, resultErr := r.db.Exec(query, friend.UserID, friend.UserID1, friend.UserID1, friend.UserID)
	if resultErr != nil {
		return resultErr
	}
	return nil
}
