package repository

import (
	"database/sql"
	"fmt"
	"projet-forum/models/entity"
	"time"
)

type UsersRepository struct {
	db *sql.DB
}

func InitUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db}
}

func (r *UsersRepository) GetAllUsers() ([]entity.User, error) {
	rows, err := r.db.Query(`
	SELECT user_id, email, username, password, bio, last_conection, image_id, salt
	FROM users;
`)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la requête : %w", err)
	}
	defer rows.Close()

	var users []entity.User

	for rows.Next() {
		var user entity.User
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

func (r *UsersRepository) Create(user entity.User) (int, error) {
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

func (r *UsersRepository) GetById(id string) (entity.User, error) {
	query := `
		SELECT email, username, bio, last_conection, image_id
		FROM users
		WHERE user_id = ?;
	`

	var user entity.User
	var lastConnection time.Time

	err := r.db.QueryRow(query, id).Scan(
		&user.Email,
		&user.Username,
		&user.Bio,
		&lastConnection,
		&user.ImageID,
	)
	if err != nil {
		return entity.User{}, fmt.Errorf("erreur lors de la requête : %w", err)
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

func (r *UsersRepository) GetByUsername(username string) (entity.User, error) {
	query := `
		SELECT user_id, email, username, password, bio, last_conection, image_id
		FROM users
		WHERE username = ?;
	`

	var user entity.User
	var lastConnection time.Time

	err := r.db.QueryRow(query, username).Scan(
		&user.UserID,
		&user.Email,
		&user.Username,
		&user.Password,
		&user.Bio,
		&lastConnection,
		&user.ImageID,
	)
	if err != nil {
		return entity.User{}, fmt.Errorf("erreur lors de la requête : %w", err)
	}

	user.LastConnection = lastConnection

	return user, nil
}

func (r *UsersRepository) GetUserByEmailOrNameAndPassword(emailOrUsername, password string) (entity.User, error) {
	query := `
		SELECT user_id, email, username, password, bio, last_conection, image_id
		FROM users
		WHERE (email = ? OR username = ?) AND password = ?;
	`

	var user entity.User
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
		return entity.User{}, fmt.Errorf("erreur lors de la requête : %w", err)
	}

	user.LastConnection = lastConnection

	return user, nil
}

func (r *UsersRepository) UpdateLastConnection(id string) error {
	query := `
		UPDATE users
		SET last_conection = ?
		WHERE user_id = ?;
	`

	_, err := r.db.Exec(query, time.Now(), id)
	if err != nil {
		return fmt.Errorf("erreur lors de la mise à jour de la dernière connexion : %w", err)
	}

	return nil
}

func (r *UsersRepository) UpdateUser(id string, user entity.User) error {
	query := `
		UPDATE users
		SET email = ?, username = ?, bio = ?, image_id = ?
		WHERE user_id = ?;
	`

	_, err := r.db.Exec(query, user.Email, user.Username, user.Bio, user.ImageID, id)
	if err != nil {
		return fmt.Errorf("erreur lors de la mise à jour de l'utilisateur : %w", err)
	}

	return nil
}

func (r *UsersRepository) UpdatePassword(id string, password string) error {
	query := `
		UPDATE users
		SET password = ?
		WHERE user_id = ?;
	`

	_, err := r.db.Exec(query, password, id)
	if err != nil {
		return fmt.Errorf("erreur lors de la mise à jour du mot de passe : %w", err)
	}

	return nil
}

func (r *UsersRepository) Delete(id string) error {
	query := `
		DELETE FROM users
		WHERE user_id = ?;
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("erreur lors de la suppression de l'utilisateur : %w", err)
	}

	return nil
}

func (r *UsersRepository) GetImageIdByUserId(id string) (string, error) {
	query := `
		SELECT image_id
		FROM users
		WHERE user_id = ?;
	`

	var imageId string

	err := r.db.QueryRow(query, id).Scan(&imageId)
	if err != nil {
		return "", fmt.Errorf("erreur lors de la requête : %w", err)
	}

	return imageId, nil
}

func (r *UsersRepository) UpdateImageIdByUserId(id string, imageId string) error {
	query := `
		UPDATE users
		SET image_id = ?
		WHERE user_id = ?;
	`

	_, err := r.db.Exec(query, imageId, id)
	if err != nil {
		return fmt.Errorf("erreur lors de la mise à jour de l'image : %w", err)
	}

	return nil
}

func (r *UsersRepository) GetIdByUsername(username string) (int, error) {
	query := `
		SELECT user_id
		FROM users
		WHERE username = ?;
	`

	var id int

	err := r.db.QueryRow(query, username).Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("erreur lors de la requête : %w", err)
	}

	return id, nil
}
