package repository

import (
	"database/sql"
	"fmt"
	"projet-forum/models"
	"time"
)

type UsersRepository struct {
	db *sql.DB
}

func InitUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db}
}

func (r *UsersRepository) getAllUsers() ([]models.User, error) {
	rows, err := r.db.Query(`
	SELECT user_id, email, username, password, bio, last_conection, image_id, salt
	FROM users;
`)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la requête : %w", err)
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		var lastConnection time.Time

		err := rows.Scan(
			&user.UserID,
			&user.Email,
			&user.Username,
			&user.Password,
			&user.Bio,
			&lastConnection,
			&user.ImageID,
			&user.Salt,
		)
		if err != nil {
			return nil, fmt.Errorf("erreur lors du scan : %w", err)
		}

		user.LastConnection = lastConnection

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erreur après l'itération : %w", err)
	}

	return users, nil
}

func (r *UsersRepository)CreateUser (user models.User) (int, error) {
	query := `
		INSERT INTO users (
			email, username, password, number_of_message_send,
			number_of_channel_create, bio, last_conection, image_id
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?);
	`

	result, err := r.db.Exec(
		query,
		user.Email,
		user.Username,
		user.Password,
		user.Bio,
		user.LastConnection,
		user.ImageID,
	)
	if err != nil {
		return -1, fmt.Errorf("erreur insertion utilisateur : %w", err)
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("erreur récupération ID : %w", err)
	}

	return int(insertedID), nil
}
