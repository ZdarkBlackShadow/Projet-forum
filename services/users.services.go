// Package services implements the business logic layer of the application.
// It contains all the service implementations that coordinate work between
// repositories and handle the core application logic.
package services

import (
        "database/sql"
        "fmt"
        "io"
        "os"
        "path/filepath"
        "projet-forum/models/dto"
        "projet-forum/models/entity"
        "projet-forum/repository"
        "projet-forum/utils"
        "time"
)

// UsersServices handles all business logic related to user management.
// It coordinates operations between the user repository and image repository
// to manage user accounts and their associated profile images.
type UsersServices struct {
        usersRepo  *repository.UsersRepository
        imagesRepo *repository.ImageRepository
}

// InitUsersServices creates a new UsersServices instance with the provided database connection.
// It initializes both the users and images repositories needed for user management.
func InitUsersServices(db *sql.DB) *UsersServices {
        return &UsersServices{
                usersRepo:  repository.InitUsersRepository(db),
                imagesRepo: repository.InitImageRepository(db),
        }
}

// Create handles the creation of a new user account with optional profile image.
// It performs the following operations:
// - Validates required user data
// - Handles profile image upload if provided
// - Hashes the user's password with a salt
// - Creates the user record in the database
// Returns the new user's ID and any error that occurred during the process.
func (s *UsersServices) Create(user entity.User, image entity.UserImage) (int, error) {
        if user.Username == "" || user.Password == "" {
                return -1, fmt.Errorf("Erreur ajout user - Données manquantes ou invalides")
        }

        if image.File != nil && image.Handler != nil {
                err := os.MkdirAll("images/users", os.ModePerm)
                if err != nil {
                        return -1, err
                }
                imageName := image.Handler.Filename
                dstPath := filepath.Join("images/users", imageName)

                dst, dstErr := os.Create(dstPath)
                if dstErr != nil {
                        return -1, dstErr
                }
                defer dst.Close()

                _, copyErr := io.Copy(dst, image.File)
                if copyErr != nil {
                        return -1, copyErr
                }

                imageId, imageErr := s.imagesRepo.Create(dstPath)
                if imageErr != nil {
                        return -1, imageErr
                }
                user.ImageID = imageId
        } else {
                user.ImageID = 14
        }

        hashedPassword, salt, passErr := utils.HashPassword(user.Password)
        if passErr != nil {
                return -1, passErr
        }

        user.Password = hashedPassword
        user.Salt = salt

        user.LastConnection = time.Now()

        userId, userErr := s.usersRepo.Create(user)
        if userErr != nil {
                return -1, userErr
        }
        return userId, nil
}

// Connect authenticates a user using their email/username and password.
// It retrieves the user's salt, hashes the provided password with it,
// and verifies the credentials against the stored user data.
// Returns the user entity if authentication is successful.
func (s *UsersServices) Connect(nameOrMail string, password string) (entity.User, error) {
        userSalt, saltErr := s.usersRepo.GetSaltByEmailOrUsername(nameOrMail)
        if saltErr != nil {
                return entity.User{}, saltErr
        }

        hashedPassword, passErr := utils.HashPasswordWithSalt(password, userSalt)
        if passErr != nil {
                return entity.User{}, passErr
        }
        user, userErr := s.usersRepo.GetUserByEmailOrNameAndPassword(nameOrMail, hashedPassword)
        if userErr != nil {
                return entity.User{}, userErr
        }

        return user, nil
}

// GetUser retrieves a user's information by their username.
// Returns the user entity if found, or an error if the user doesn't exist.
func (s *UsersServices) GetUser(username string) (entity.User, error) {
        user, userErr := s.usersRepo.GetByUsername(username)
        if userErr != nil {
                return entity.User{}, userErr
        }
        return user, nil
}

// GetAllInformationAboutOneUser retrieves complete user information including their profile image
// using the provided JWT token for authentication.
// Returns a UserInformation DTO containing the user's details and profile image path.
func (s *UsersServices) GetAllInformationAboutOneUser(token string) (dto.UserInformation, error) {
        userId, jwtErr := utils.VerifyJWT(token)
        if jwtErr != nil {
                return dto.UserInformation{}, jwtErr
        }

        user, err := s.usersRepo.GetById(userId)
        if err != nil {
                return dto.UserInformation{}, err
        }

        pathImage, err := s.imagesRepo.GetById(user.ImageID)
        if err != nil {
                return dto.UserInformation{}, err
        }

        return dto.UserInformation{
                Username: user.Username,
                Email:    user.Email,
                Bio:      user.Bio,
                Image:    pathImage,
        }, nil
}
