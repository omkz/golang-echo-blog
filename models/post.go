package models

import (
	_"database/sql"
	"database/sql"
)

type Post struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
}

var con *sql.DB

func PostAll() ([]*Post, error) {

	con :=con.CreateCon()

	rows, err := con.Query("select * from posts")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := []*Post{}

	for rows.Next() {
		post := &Post{}
		err := rows.Scan(&post.Id, &post.Title, &post.Content)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
