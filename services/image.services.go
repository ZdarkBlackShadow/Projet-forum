package services

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"projet-forum/repository"
	"projet-forum/utils"
)

type ImageServices struct {
	imagesRepo *repository.ImageRepository
}

func InitImageServices(db *sql.DB) *ImageServices {
	return &ImageServices{
		imagesRepo: repository.InitImageRepository(db),
	}
}

func (s *ImageServices) GetImageByName(token string, filename string) ([]byte, string, error) {
	userId, jwtErr := utils.VerifyJWT(token)
	if jwtErr != nil {
		return nil, "", jwtErr
	}

	if userId != "0" {
		
	} //rajouter les permisions users

	filepath := fmt.Sprintf("images/users/%s", filename)
	file, fileErr := os.Open(filepath)
	if fileErr != nil {
		return nil, "", fileErr
	}
	defer file.Close()

	header := make([]byte, 512)
	n, err := file.Read(header)
	if err != nil && err != io.EOF {
		return nil, "", err
	}

	contentType := http.DetectContentType(header[:n])

	rest, err := io.ReadAll(file)
	if err != nil {
		return nil, "", err
	}

	imageData := append(header[:n], rest...)

	return imageData, contentType, nil
}
