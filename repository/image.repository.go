package repository

import "database/sql"

type ImageRepository struct {
	db *sql.DB
}

func InitImageRepository(db *sql.DB) *ImageRepository {
	return &ImageRepository{db}
}

func (r *ImageRepository) Create(path string) (int, error) {
	/*
	* test
	*/
	query := "INSERT INTO `image` (`path`) VALUES (?);"

	result, resultErr := r.db.Exec(query, path)
	if resultErr != nil {
		return -1, resultErr
	}

	insertId, insertIdErr := result.LastInsertId()
	if insertIdErr != nil {
		return -1, insertIdErr
	}
	return int(insertId), nil
}

func (r *ImageRepository) GetById(id int) (string, error) {
	query := "SELECT `path` FROM `image` WHERE `image_id` = ?;"

	var path string
	err := r.db.QueryRow(query, id).Scan(&path)
	if err != nil {
		return "", err
	}
	return path, nil
}

func (r *ImageRepository) CanAccess(imageId int, userId int) (bool, error) {
	query := "SELECT `image_id` FROM `image_than_user_can_access` WHERE `image_id` = ? AND `user_id` = ?;"

	var id int
	err := r.db.QueryRow(query, imageId, userId).Scan(&id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *ImageRepository) GiveAccess(imageId int, userId int) error {
	query := "INSERT INTO `image_than_user_can_access` (`image_id`, `user_id`) VALUES (?, ?);"

	_, err := r.db.Exec(query, imageId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *ImageRepository) RemoveAccess(imageId int, userId int) error {
	query := "DELETE FROM `image_than_user_can_access` WHERE `image_id` = ? AND `user_id` = ?;"

	_, err := r.db.Exec(query, imageId, userId)
	if err != nil {
		return err
	}
	return nil
}
