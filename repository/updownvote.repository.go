package repository

import "database/sql"

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
