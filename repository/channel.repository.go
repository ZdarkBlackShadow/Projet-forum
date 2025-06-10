package repository

import (
	"database/sql"
	"projet-forum/models/entity"
	"time"
)

type ChannelRepository struct {
	db *sql.DB
}

func InitChannelRepository(db *sql.DB) *ChannelRepository {
	return &ChannelRepository{db}
}

func (r *ChannelRepository) CreateChannel(channelInfo entity.Channel) (int, error) {
	query := "INSERT INTO channels (name, description, created_at, private, image_id, state_id, user_id) VALUES (?, ?, ?, ?, ?, ?, ?)"

	result, resultErr := r.db.Exec(query, channelInfo.Name, channelInfo.Description, channelInfo.CreatedAt, channelInfo.Private, channelInfo.ImageID, channelInfo.StateID, channelInfo.UserID)
	if resultErr != nil {
		return -1, resultErr
	}

	channelId, channelIdErr := result.LastInsertId()
	if channelIdErr != nil {
		return -1, channelIdErr
	}
	return int(channelId), nil
}

func (r *ChannelRepository) GetChannelById(channelId int) (entity.Channel, error) {
	query := "SELECT channel_id, name, description,	created_at,	private	image_id, state_id,	user_id FROM channels WHERE id = ?"

	var channel entity.Channel
	err := r.db.QueryRow(query, channelId).Scan(&channel.ChannelID, &channel.Name, &channel.Description, &channel.CreatedAt, &channel.Private, &channel.ImageID, &channel.StateID, &channel.UserID)
	if err != nil {
		return entity.Channel{}, err
	}

	return channel, nil
}

func (r *ChannelRepository) GetChannelByName(channelName string) (entity.Channel, error) {
	query := "SELECT channel_id, name, description,	created_at,	private	image_id, state_id,	user_id FROM channels WHERE name = ?"

	var channel entity.Channel
	err := r.db.QueryRow(query, channelName).Scan(&channel.ChannelID, &channel.Name, &channel.Description, &channel.CreatedAt, &channel.Private, &channel.ImageID, &channel.StateID, &channel.UserID)
	if err != nil {
		return entity.Channel{}, err
	}

	return channel, nil
}

func (r *ChannelRepository) GetChannelsByUserId(userId int) ([]entity.Channel, error) {
	query := "SELECT channel_id, name, description,	created_at,	private	image_id, state_id,	user_id FROM channels WHERE user_id = ?"

	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var channels []entity.Channel
	for rows.Next() {
		var channel entity.Channel
		err := rows.Scan(&channel.ChannelID, &channel.Name, &channel.Description, &channel.CreatedAt, &channel.Private, &channel.ImageID, &channel.StateID, &channel.UserID)
		if err != nil {
			return nil, err
		}
		channels = append(channels, channel)
	}

	return channels, nil
}

func (r *ChannelRepository) UpdateChannel(channelInfo entity.Channel) error {
	query := "UPDATE channels SET name = ?, description = ?, created_at = ?, private = ?, image_id = ?, state_id = ?, user_id = ? WHERE channel_id = ?"

	_, err := r.db.Exec(query, channelInfo.Name, channelInfo.Description, channelInfo.CreatedAt, channelInfo.Private, channelInfo.ImageID, channelInfo.StateID, channelInfo.UserID, channelInfo.ChannelID)
	if err != nil {
		return err
	}

	return nil
}

func (r *ChannelRepository) DeleteChannel(channelId int) error {
	query := "DELETE FROM channels WHERE channel_id = ?"

	_, err := r.db.Exec(query, channelId)
	if err != nil {
		return err
	}

	return nil
}

func (r *ChannelRepository) VerifyAccess(channelId int, userId int) (bool, error) {
	query := "SELECT COUNT(*) FROM channels WHERE channel_id = ? AND user_id = ?"

	var count int
	err := r.db.QueryRow(query, channelId, userId).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *ChannelRepository) GetStateIdByName(name string) int {
	query := "SELECT state_id FROM state WHERE name = ?"

	var stateId int
	err := r.db.QueryRow(query, name).Scan(&stateId)
	if err != nil {
		return -1
	}
	return stateId
}

func (r *ChannelRepository) CreateChannelInvitation(user_id_creator int, user_id_invite int, channelId int) error {
	query := "INSERT INTO channel_invitation (user_id_creator, user_id_invite, channel_id, created_at) VALUES (?, ?, ?, ?)"

	_, err := r.db.Exec(query, user_id_creator, user_id_invite, channelId, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (r *ChannelRepository) GetAllChannelInvitations(user_id_invite int) ([]entity.ChannelInvitation, error) {
	query := "SELECT user_id_creator, user_id_invite, channel_id, created_at FROM channel_invitation WHERE user_id_invite = ?;"

	rows, err := r.db.Query(query, user_id_invite)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invitations []entity.ChannelInvitation

	for rows.Next() {
		var inv entity.ChannelInvitation
		err := rows.Scan(&inv.UserID, &inv.UserID1, &inv.ChannelID, &inv.CreatedAt)
		if err != nil {
			return nil, err
		}
		invitations = append(invitations, inv)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return invitations, nil
}

func (r *ChannelRepository) DeleteChannelInvitation(user_id_creator int, user_id_invite int, channelId int) error {
	query := "DELETE FROM channel_invitation WHERE user_id_creator = ? AND user_id_invite = ? AND channel_id = ?"

	_, err := r.db.Exec(query,user_id_creator, user_id_invite, channelId)
	if err != nil {
		return err
	}
	return nil
}

func (r *ChannelRepository) GetChannelInvitation(user_id_creator int, user_id_invite int, channelId int) (entity.ChannelInvitation, error) {
	query := "SELECT user_id_creator, user_id_invite, channel_id, created_at FROM channel_invitation WHERE user_id_creator = ? AND user_id_invite = ? AND channel_id = ?"

	var inv entity.ChannelInvitation
	err := r.db.QueryRow(query, user_id_creator, user_id_invite, channelId).Scan(&inv.UserID, &inv.UserID1, &inv.ChannelID, &inv.CreatedAt)
	if err != nil {
		return entity.ChannelInvitation{}, err
	}
	return inv, nil
}

func (r *ChannelRepository) AddUserToChannel(user_id int, channelId int) error {
	query := "INSERT INTO users_who_can_acces (user_id, channel_id) VALUES (?, ?)"

	_, err := r.db.Exec(query, user_id, channelId)
	if err != nil {
		return err
	}
	return nil
}

func (r *ChannelRepository) RemoveUserFromChannel(user_id int, channelId int) error {
	query := "DELETE FROM users_who_can_acces WHERE user_id = ? AND channel_id = ?"

	_, err := r.db.Exec(query, user_id, channelId)
	if err != nil {
		return err
	}
	return nil
}

func (r *ChannelRepository) GetChannelUsers(channelId int) ([]entity.User, error) {
	query := "SELECT user_id FROM users_who_can_acces WHERE channel_id = ?"

	rows, err := r.db.Query(query, channelId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User

	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.UserID)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *ChannelRepository) GetChannelUsersCount(channelId int) (int, error) {
	query := "SELECT COUNT(*) FROM users_who_can_acces WHERE channel_id = ?"

	var count int
	err := r.db.QueryRow(query, channelId).Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (r *ChannelRepository) GetAllPublicChannels() ([]entity.Channel, error) {
	query := "SELECT channel_id, name, description, created_at, private, image_id, state_id, user_id FROM channels WHERE private = 0"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var channels []entity.Channel

	for rows.Next() {
		var channel entity.Channel
		err := rows.Scan(&channel.ChannelID, &channel.Name, &channel.Description, &channel.CreatedAt, &channel.Private, &channel.ImageID, &channel.StateID, &channel.UserID)
		if err != nil {
			return nil, err
		}
		channels = append(channels, channel)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return channels, nil
}
