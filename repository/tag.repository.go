package repository

import (
	"database/sql"
	"time"
)

type TagRepository struct {
	db *sql.DB
}

func InitTagRepository(db *sql.DB) *TagRepository {
	return &TagRepository{db}
}

func (r *TagRepository) CreateTag(tagName string) (int, error) {
	query := "INSERT INTO tags (name, created_at) VALUES (?, ?)"

	result, resultErr := r.db.Exec(query, tagName, time.Now())
	if resultErr != nil {
		return -1, resultErr
	}

	tagId, tagIdErr := result.LastInsertId()
	if tagIdErr != nil {
		return -1, tagIdErr
	}
	return int(tagId), nil
}

func (r *TagRepository) GetTagIdByName(tagName string) (int, error) {
	query := "SELECT tag_id FROM tags WHERE name = ?"

	var tagId int
	err := r.db.QueryRow(query, tagName).Scan(&tagId)
	if err != nil {
		return -1, err
	}
	return tagId, nil
}

func (r *TagRepository) DeleteTag(tagId int) error {
	query := "DELETE FROM tags WHERE tag_id = ?"

	_, err := r.db.Exec(query, tagId)
	if err != nil {
		return err
	}

	return nil
}

func (r *TagRepository) AddTagToChannel (channelId int, tagId int) error {
	query := "INSERT INTO channel_tags (channel_id, tag_id) VALUES (?, ?)"

	_, err := r.db.Exec(query, channelId, tagId)
	if err != nil {
		return err
	}

	return nil
}

func (r *TagRepository) RemoveTagFromChannel (channelId int, tagId int) error {
	query := "DELETE FROM channel_tags WHERE channel_id = ? AND tag_id = ?"

	_, err := r.db.Exec(query, channelId, tagId)
	if err != nil {
		return err
	}

	return nil
}

func (r *TagRepository) GetTagsByChannelId(channelId int) ([]string, error) {
	query := "SELECT tags.name FROM tags INNER JOIN channel_tags ON tags.tag_id = channel_tags.tag_id WHERE channel_tags.channel_id = ?"

	rows, err := r.db.Query(query, channelId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []string
	for rows.Next() {
		var tag string
		err := rows.Scan(&tag)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}
