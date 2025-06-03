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

func (r *UsersRepository)GetAllUsers() ([]models.User, error) {
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

func (r *UsersRepository)Create (user models.User) (int, error) {
	query := `
		INSERT INTO users (
			email, username, password, 
			bio, last_conection, image_id, salt
		)
		VALUES (?, ?, ?, ?, ?, ?, ?);
	`

	result, err := r.db.Exec(
		query,
		user.Email,
		user.Username,
		user.Password,
		user.Bio,
		user.LastConnection,
		user.ImageID,
		user.Salt,
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

func (r *UsersRepository)GetById(id string) (models.User, error) {
	query := `
		SELECT email, username, bio, last_conection, image_id
		FROM users
		WHERE user_id = ?;
	`

	var user models.User
	var lastConnection time.Time

	err := r.db.QueryRow(query, id).Scan(
		&user.Email,
		&user.Username,
		&user.Bio,
		&lastConnection,
		&user.ImageID,
	)
	if err != nil {
		return models.User{}, fmt.Errorf("erreur lors de la requête : %w", err)
	}

	user.LastConnection = lastConnection

	return user, nil
}

func (r *UsersRepository) GetSaltByEmailOrUsername(emailOrUsername string) (string, error) {
	query := `
		SELECT salt
		FROM users
		WHERE email = ? OR username = ?;
	`

	var salt string

	err := r.db.QueryRow(query, emailOrUsername, emailOrUsername).Scan(&salt)
	if err != nil {
		return "", fmt.Errorf("erreur lors de la requête : %w", err)
	}

	return salt, nil
}

func (r *UsersRepository) GetUserByEmailOrNameAndPassword(emailOrUsername, password string) (models.User, error) {
	query := `
		SELECT user_id, email, username, password, bio, last_conection, image_id
		FROM users
		WHERE (email = ? OR username = ?) AND password = ?;
	`

	var user models.User
	var lastConnection time.Time

	err := r.db.QueryRow(query, emailOrUsername, emailOrUsername, password).Scan(
		&user.UserID,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Bio,
		&lastConnection,
		&user.ImageID,
	)
	if err != nil {
		return models.User{}, fmt.Errorf("erreur lors de la requête : %w", err)
	}

	user.LastConnection = lastConnection

	return user, nil
}