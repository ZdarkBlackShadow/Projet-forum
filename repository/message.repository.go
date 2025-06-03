package repository

import (
	"database/sql"
	"projet-forum/models"
	"time"
)

type MessageRepository struct {
	db *sql.DB
}

func InitMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{db}
}

func (r *MessageRepository) CreateMessage(message models.Message) (int, error) {
	result, err := r.db.Exec("INSERT INTO `message`(`text`, `created_at`, `edited`, `image`, `user_id`, `channel_id`) VALUES ('?','?','?','?','?','?');",
		message.Text,
		time.Now(),
		false,
		message.Image,
		message.UserID,
		message.Channel_id,
	)
	if err != nil {
		return -1, err
	}
	lastInserted, insertedErr := result.LastInsertId()
	if insertedErr != nil {
		return -1, insertedErr
	}
	return int(lastInserted), nil
}

func (r *MessageRepository) GetMessageByID(id int) (models.Message, error) {
	var message models.Message
	row := r.db.QueryRow("SELECT `id`, `text`, `created_at`, `edited`, `image`, `user_id`, `channel_id` FROM `message` WHERE `id` = ?;", id)
	err := row.Scan(&message.MessageTextID, &message.Text, &message.CreatedAt, &message.Edited, &message.Image, &message.UserID, &message.Channel_id)
	if err != nil {
		return models.Message{}, err
	}
	return message, nil
}

func (r *MessageRepository) GetMessagesByChannelID(channelID int) ([]models.Message, error) {
	var messages []models.Message
	rows, err := r.db.Query("SELECT `id`, `text`, `created_at`, `edited`, `image`, `user_id`, `channel_id` FROM `message` WHERE `channel_id` = ?;", channelID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var message models.Message
		err := rows.Scan(&message.MessageTextID, &message.Text, &message.CreatedAt, &message.Edited, &message.Image, &message.UserID, &message.Channel_id)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func (r *MessageRepository) UpdateMessage(message models.Message) error {
	_, err := r.db.Exec("UPDATE `message` SET `text` = '?', `edited` = '?', `image` = '?' WHERE `id` = '?';",
		message.Text,
		true,
		message.Image,
		message.MessageTextID,
	)
	if err != nil {
		return err
	}
	return nil
}