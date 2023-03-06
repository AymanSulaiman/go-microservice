package controller

import (
	"database/sql"
	"errors"
	"go-microservice/database"
)

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func GetPost(id int) (*Post, error) {
	var post Post

	err := database.DB.QueryRow("SELECT id, title, content FROM posts WHERE id = ?", id).Scan(&post.ID, &post.Title, &post.Content)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("Post not found")
		}
		return nil, err
	}

	return &post, nil
}

func CreatePost(post *Post) (int, error) {
	result, err := database.DB.Exec("INSERT INTO posts (title, content) VALUES (?, ?)", post.Title, post.Content)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
