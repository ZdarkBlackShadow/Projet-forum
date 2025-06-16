package repository

import (
	"database/sql"
	"projet-forum/models/entity"
)

type UpDownVoteRepository struct {
	db *sql.DB
}

func InitUpDownVoteRepository(db *sql.DB) *UpDownVoteRepository {
	return &UpDownVoteRepository{db}
}

func (r *UpDownVoteRepository) Create(userId int, messageId int, upDownVote int) error {
	query := "INSERT INTO up_down (user_id, message_id, up_down_vote) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, userId, messageId, upDownVote)
	return err
}

func (r *UpDownVoteRepository) Update(userId int, messageId int, upDownVote int) error {
	query := "UPDATE up_down SET up_down_vote = ? WHERE user_id = ? AND message_id = ?"
	_, err := r.db.Exec(query, upDownVote, userId, messageId)
	return err
}

func (r *UpDownVoteRepository) GetAllUpDownVoteFromMessageId(messageId int) ([]entity.UpDown, error) {
	query := "SELECT `user_id`, `message_text_id`, `up_down_vote_id` FROM `up_down` WHERE message_text_id = ?"
	rows, err := r.db.Query(query, messageId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var upDownVotes []entity.UpDown
	for rows.Next() {
		var upDownVote entity.UpDown
		err := rows.Scan(&upDownVote.UserID, &upDownVote.MessageTextID, &upDownVote.UpDownVoteID)
		if err != nil {
			return nil, err
		}
		upDownVotes = append(upDownVotes, upDownVote)
	}

	return upDownVotes, nil
}
