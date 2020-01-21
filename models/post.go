package models

import (
	_"database/sql"
	"database/sql"
	"github.com/omkz/golang-echo-blog/db"
)

type Post struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
}

var con *sql.DB

func PostAll() ([]*Post) {

	con = db.CreateCon()

	rows, err := con.Query("select * from posts")

	if err != nil {
		return nil
	}

	defer rows.Close()

	posts := []*Post{}

	for rows.Next() {
		post := &Post{}
		err := rows.Scan(&post.Id, &post.Title, &post.Content)
		if err != nil {
			return nil
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return nil
	}

	return posts
}
