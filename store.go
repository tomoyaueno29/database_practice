package main

import (
	"database/sql"
	"errors"
)

type Post struct {
	Id int
	Content string
	Author string
	Comments []Comment
}

type Comment struct {
	Id int
	Content string
	Author string
	Post *Post
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres",
		"user=tomoyaueno dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func (comment *Comment) Create() (err error){
	if comment.Post == nil {
		err = errors.New("No posts")
		return
	}
	err = Db.QueryRow(`insert into comments (content, author, post_id)
	values ($1, $2, $3) returning id`, comment.Content, comment.Author,
	comment.Post.Id).Scan(&comment.Id)
	return
}