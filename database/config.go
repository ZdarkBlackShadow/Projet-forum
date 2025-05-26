package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"projet-forum/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init() error {
	var err error
	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PWD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pwd, host, port, name)
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
		fmt.Printf("User ID: %d, Email: %s, Username: %s, Password: %s, Bio: %s, Last Connection: %s, Image ID: %d, Salt : %s\n",
			user.UserID, user.Email, user.Username, user.Password, user.Bio, user.LastConnection, user.ImageID, user.Salt)
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
