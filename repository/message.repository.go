package repository

import (
	"database/sql"
	"projet-forum/models/entity"
	"time"
)

type MessageRepository struct {
	db *sql.DB
}

func InitMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{db}
}

func (r *MessageRepository) CreateMessage(message entity.Message) (int, error) {
	result, err := r.db.Exec("INSERT INTO `message`(`text`, `created_at`, `edited`, `image`, `user_id`, `channel_id`) VALUES (?,?,?,?,?,?);",
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

func (r *MessageRepository) GetMessageByID(id int) (entity.Message, error) {
	var message entity.Message
	row := r.db.QueryRow("SELECT `message_text_id`, `text`, `created_at`, `edited`, `image`, `user_id`, `channel_id` FROM `message` WHERE `message_text_id` = ?;", id)
	err := row.Scan(&message.MessageTextID, &message.Text, &message.CreatedAt, &message.Edited, &message.Image, &message.UserID, &message.Channel_id)
	if err != nil {
		return entity.Message{}, err
	}
	return message, nil
}

func (r *MessageRepository) GetMessagesByChannelID(channelID int) ([]entity.Message, error) {
	var messages []entity.Message
	rows, err := r.db.Query("SELECT `message_text_id`, `text`, `created_at`, `edited`, `image`, `user_id`, `channel_id` FROM `message` WHERE `channel_id` = ?;", channelID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var message entity.Message
		err := rows.Scan(&message.MessageTextID, &message.Text, &message.CreatedAt, &message.Edited, &message.Image, &message.UserID, &message.Channel_id)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func (r *MessageRepository) UpdateMessage(message entity.Message) error {
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

func (r *MessageRepository) DeleteMessage(id int) error {
	_, err := r.db.Exec("DELETE FROM `message` WHERE `id` = '?';", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepository) DeleteMessagesByChannelID(channelID int) error {
	_, err := r.db.Exec("DELETE FROM `message` WHERE `channel_id` = '?';", channelID)
	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepository) DeleteMessagesByUserID(userID int) error {
	_, err := r.db.Exec("DELETE FROM `message` WHERE `user_id` = '?';", userID)
	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepository) AddUpDownVote(messageId int, userId int, voteId int) error {
	_, err := r.db.Exec("INSERT INTO `up_down`(`user_id`, `message_text_id`, `up_down_vote_id`) VALUES (?, ?, ?);",
		userId,
		messageId,
		voteId,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *MessageRepository) UpdateUpDownVote(messageId int, userId int, new_vote int) error {
	_, err := r.db.Exec("UPDATE `up_down` SET `up_down_vote_id` = '?' WHERE `user_id` = '?' AND `message_text_id` = '?';",
		new_vote,
		userId,
		messageId,
	)
	if err != nil {
		return err
	}
	return nil
}
