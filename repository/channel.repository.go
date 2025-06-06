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

func (r *ChannelRepository) GetStateIdByName(name string) int{
    query := "SELECT state_id FROM state WHERE name = ?"

    var stateId int
    err := r.db.QueryRow(query, name).Scan(&stateId)
    if err != nil {
        return -1
    }
    return stateId
}

func (r *ChannelRepository) CreateTag(tagName string) (int, error) {
	query := "INSERT INTO tags (name, created_at) VALUES (?, ?)"

	result, resultErr := r.db.Exec(query, tagName, time.Now())
	if resultErr != nil {
		return -1, resultErr
	}

	tagId, tagIdErr := result.LastInsertId()
	if tagIdErr != nil {
		return -1, tagIdErr
	}
	return int(tagId), nil
}

func (r *ChannelRepository) GetTagIdByName(tagName string) int {
	query := "SELECT tag_id FROM tags WHERE name = ?"

	var tagId int
	err := r.db.QueryRow(query, tagName).Scan(&tagId)
	if err != nil {
		return -1
	}
	return tagId
}

func (r *ChannelRepository) DeleteTag(tagId int) error {
	query := "DELETE FROM tags WHERE tag_id = ?"

	_, err := r.db.Exec(query, tagId)
	if err != nil {
		return err
	}

	return nil
}