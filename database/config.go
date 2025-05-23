package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"projet-forum/models"
	"projet-forum/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init() error {
	var err error
	utils.LoadEnvFile(".env")
	databaseName := os.Getenv("DATABASE_NAME")
	fmt.Println(databaseName)
	dsn := "root:@tcp(127.0.0.1:3306)/" + databaseName + "?parseTime=true"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	fmt.Println("Connexion réussie à la base de données")
	users, err := GetAllUser()
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Printf("User ID: %d, Email: %s, Username: %s, Password: %s, Bio: %s, Last Connection: %s, Image ID: %d\n",
			user.UserID, user.Email, user.Username, user.Password, user.Bio, user.LastConnection, user.ImageID)
	}
	return nil
}

func Close() {
	if db != nil {
		db.Close()
	}
}

func GetAllUser() ([]models.User, error) {
	rows, err := db.Query(`
	SELECT user_id, email, username, password, bio, last_conection, image_id
	FROM users;
`)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la requête : %w", err)
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		var bio sql.NullString
		var lastConnection time.Time

		err := rows.Scan(
			&user.UserID,
			&user.Email,
			&user.Username,
			&user.Password,
			&bio,
			&lastConnection,
			&user.ImageID,
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
