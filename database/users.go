package database

import (
	"fmt"
	"projet-forum/models"
)

func InsertUser(user models.User) (int64, error) {
	query := `
		INSERT INTO users (
			email, username, password, number_of_message_send,
			number_of_channel_create, bio, last_conection, image_id
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?);
	`

	result, err := db.Exec(
		query,
		user.Email,
		user.Username,
		user.Password,
		user.Bio,
		user.LastConnection,
		user.ImageID,
	)
	if err != nil {
		return 0, fmt.Errorf("erreur insertion utilisateur : %w", err)
	}

	insertedID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("erreur récupération ID : %w", err)
	}

	return insertedID, nil
}
